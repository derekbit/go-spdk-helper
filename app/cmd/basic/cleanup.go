package basic

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	commontypes "github.com/longhorn/go-common-libs/types"

	"github.com/longhorn/go-spdk-helper/pkg/initiator"
	"github.com/longhorn/go-spdk-helper/pkg/types"
	"github.com/longhorn/go-spdk-helper/pkg/util"
)

func CleanupLocalV2DevicesCmd() cli.Command {
	return cli.Command{
		Name: "cleanup-local-v2-devices",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host-proc",
				Usage: fmt.Sprintf("The host proc path of namespace executor. By default %v", commontypes.HostProcDirectory),
				Value: commontypes.HostProcDirectory,
			},
		},
		Usage: "Clean up local Longhorn v2 (NVMe/TCP backed) device-mapper devices and endpoints whose frontend initiator lives on this node. Intended for the v2 instance-manager pre-stop hook.",
		Action: func(c *cli.Context) {
			if err := cleanupLocalV2Devices(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run cleanup local v2 devices command")
			}
		},
	}
}

// cleanupLocalV2Devices tears down the device-mapper device, endpoint and NVMe/TCP
// connection for every Longhorn v2 volume whose frontend initiator is connected on
// this node. Only subsystems carrying the Longhorn NQN prefix are touched, so v1
// volumes (backed by iSCSI/dm-crypt, not NVMe) are never affected.
func cleanupLocalV2Devices(c *cli.Context) error {
	hostProc := c.String("host-proc")

	executor, err := util.NewExecutor(hostProc)
	if err != nil {
		return err
	}

	subsystems, err := initiator.GetSubsystems(executor)
	if err != nil {
		return fmt.Errorf("failed to list NVMe subsystems: %w", err)
	}

	nqnPrefix := types.NQNPrefix + ":"
	for _, subsystem := range subsystems {
		if !strings.HasPrefix(subsystem.NQN, nqnPrefix) || subsystem.NQN == types.InternalHostNQN {
			continue
		}

		name := strings.TrimPrefix(subsystem.NQN, nqnPrefix)
		log := logrus.WithFields(logrus.Fields{"name": name, "nqn": subsystem.NQN})

		i, err := initiator.NewInitiator(name, hostProc, &initiator.NVMeTCPInfo{SubsystemNQN: subsystem.NQN}, nil)
		if err != nil {
			log.WithError(err).Warn("Failed to create initiator for local v2 device cleanup, skipping")
			continue
		}

		// Best-effort per volume: remove the linear dm device and endpoint, then
		// disconnect the NVMe/TCP target. Do not abort the whole cleanup on a
		// single volume failure so the remaining devices are still cleaned up.
		if _, err := i.Stop(nil, true, false, false); err != nil {
			log.WithError(err).Warn("Failed to clean up local v2 device, continuing with the rest")
			continue
		}
		log.Info("Cleaned up local v2 device")
	}

	return nil
}
