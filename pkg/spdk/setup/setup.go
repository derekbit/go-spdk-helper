package setup

import (
	"fmt"

	commonNs "github.com/longhorn/go-common-libs/ns"

	"github.com/longhorn/go-spdk-helper/pkg/types"
)

const (
	spdkSetupPath = "/usr/src/spdk/scripts/setup.sh"
)

func Bind(deviceAddr, deviceDriver string, executor *commonNs.Executor) (string, error) {
	envs := []string{
		fmt.Sprintf("%s=%s", "PCI_ALLOWED", deviceAddr),
		fmt.Sprintf("%s=%s", "DRIVER_OVERRIDE", deviceDriver),
	}

	cmdArgs := []string{
		spdkSetupPath,
		"bind",
	}

	outputStr, err := executor.Execute(envs, "bash", cmdArgs, types.ExecuteTimeout)
	if err != nil {
		return "", err
	}

	return outputStr, nil
}
