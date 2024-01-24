package spdksetup

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	commonTypes "github.com/longhorn/go-common-libs/types"

	"github.com/longhorn/go-spdk-helper/pkg/util"

	spdksetup "github.com/longhorn/go-spdk-helper/pkg/spdk/setup"
)

func Cmd() cli.Command {
	return cli.Command{
		Name: "spdk-setup",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host-proc",
				Usage: fmt.Sprintf("The host proc path of namespace executor. By default %v", commonTypes.ProcDirectory),
				Value: commonTypes.ProcDirectory,
			},
		},
		Subcommands: []cli.Command{
			BindCmd(),
		},
	}
}

func BindCmd() cli.Command {
	return cli.Command{
		Name: "bind",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "device-addr",
				Usage:    "The device address to bind",
				Required: true,
			},
			cli.StringFlag{
				Name:     "device-driver",
				Usage:    "The device driver to bind to",
				Required: true,
			},
			cli.StringFlag{ // Add this line for host-proc in bind subcommand
				Name:  "host-proc",
				Usage: "The host proc path of namespace executor. By default /proc",
				Value: commonTypes.ProcDirectory,
			},
		},
		Usage: "Bind the device to SPDK",
		Action: func(c *cli.Context) {
			if err := bind(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to bind device %v to SPDK", c.String("device-addr"))
			}
		},
	}
}

func bind(c *cli.Context) error {
	executor, err := util.NewExecutor(c.GlobalString("host-proc"))
	if err != nil {
		return err
	}

	out, err := spdksetup.Bind(c.String("device-addr"), c.String("device-driver"), executor)
	if err != nil {
		return err
	}

	return util.PrintObject(out)
}
