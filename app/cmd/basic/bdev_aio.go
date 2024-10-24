package basic

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/longhorn/go-spdk-helper/pkg/spdk/client"
	"github.com/longhorn/go-spdk-helper/pkg/util"
)

func BdevAioCmd() cli.Command {
	return cli.Command{
		Name:      "bdev-aio",
		ShortName: "aio",
		Subcommands: []cli.Command{
			BdevAioCreateCmd(),
			BdevAioDeleteCmd(),
			BdevAioGetCmd(),
		},
	}
}

func BdevAioCreateCmd() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "create a bdev aio based on a block device: create --file-path <BLOCK DEVICE PATH> --bdev-name <BDEV NAME> --block-size <BLOCK SIZE>",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "file-path, f",
				Usage:    "Path to device or file",
				Required: true,
			},
			cli.StringFlag{
				Name:     "bdev-name, n",
				Usage:    "Bdev name to use",
				Required: true,
			},
			cli.Uint64Flag{
				Name:  "block-size, b",
				Usage: "The block size in bytes. By default 4096",
				Value: 4096,
			},
		},
		Action: func(c *cli.Context) {
			if err := bdevAioCreate(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run create bdev aio command")
			}
		},
	}
}

func bdevAioCreate(c *cli.Context) error {
	spdkCli, err := client.NewClient()
	if err != nil {
		return err
	}

	bdevName, err := spdkCli.BdevAioCreate(c.String("file-path"), c.String("bdev-name"), c.Uint64("block-size"))
	if err != nil {
		return err
	}

	return util.PrintObject(map[string]string{"bdev_name": bdevName})
}

func BdevAioDeleteCmd() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "delete a bdev aio using a block device: delete <BDEV NAME>",
		Action: func(c *cli.Context) {
			if err := bdevAioDelete(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run delete bdev aio command")
			}
		},
	}
}

func bdevAioDelete(c *cli.Context) error {
	spdkCli, err := client.NewClient()
	if err != nil {
		return err
	}

	deleted, err := spdkCli.BdevAioDelete(c.Args().First())
	if err != nil {
		return err
	}

	return util.PrintObject(deleted)
}

func BdevAioGetCmd() cli.Command {
	return cli.Command{
		Name: "get",
		Flags: []cli.Flag{
			cli.Uint64Flag{
				Name:  "timeout, t",
				Usage: "Determine the timeout of the execution",
				Value: 0,
			},
		},
		Usage: "get all AIO bdevs if a bdev name is not specified: \"get\", or \"get <AIO BDEV NAME>\"",
		Action: func(c *cli.Context) {
			if err := bdevAioGet(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run get bdev aio command")
			}
		},
	}
}

func bdevAioGet(c *cli.Context) error {
	spdkCli, err := client.NewClient()
	if err != nil {
		return err
	}

	bdevAioGetResp, err := spdkCli.BdevAioGet(c.Args().First(), 0)
	if err != nil {
		return err
	}

	return util.PrintObject(bdevAioGetResp)
}
