package basic

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/longhorn/go-spdk-helper/pkg/spdk/client"
	"github.com/longhorn/go-spdk-helper/pkg/types"
	"github.com/longhorn/go-spdk-helper/pkg/util"
)

func BdevLvstoreCmd() cli.Command {
	return cli.Command{
		Name:      "bdev-lvstore",
		ShortName: "lvs",
		Subcommands: []cli.Command{
			BdevLvstoreCreateCmd(),
			BdevLvstoreDeleteCmd(),
			BdevLvstoreGetCmd(),
			BdevLvstoreRenameCmd(),
			BdevLvstoreGetLvolsCmd(),
		},
	}
}

func BdevLvstoreCreateCmd() cli.Command {
	return cli.Command{
		Name: "create",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "bdev-name",
				Usage:    "The bdev on which to construct logical volume store",
				Required: true,
			},
			cli.StringFlag{
				Name:     "lvs-name",
				Usage:    "Name of the logical volume store to create",
				Required: true,
			},
			cli.UintFlag{
				Name:  "cluster-size",
				Usage: "Logical volume store cluster size, by default 1MiB",
				Value: types.MiB,
			},
		},
		Usage: "create a bdev lvstore based on a block device: \"create --bdev-name <BDEV NAME> --lvs-name <LVSTORE NAME>\"",
		Action: func(c *cli.Context) {
			if err := bdevLvstoreCreate(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run create bdev lvstore command")
			}
		},
	}
}

func bdevLvstoreCreate(c *cli.Context) error {
	spdkCli, err := client.NewClient()
	if err != nil {
		return err
	}

	uuid, err := spdkCli.BdevLvolCreateLvstore(c.String("bdev-name"), c.String("lvs-name"), uint32(c.Uint("cluster-size")))
	if err != nil {
		return err
	}

	return util.PrintObject(uuid)
}

func BdevLvstoreRenameCmd() cli.Command {
	return cli.Command{
		Name: "rename",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "old-name",
				Usage:    "Old name of the logical volume store",
				Required: true,
			},
			cli.StringFlag{
				Name:     "new-name",
				Usage:    "New name of the logical volume store",
				Required: true,
			},
		},
		Usage: "rename a bdev lvstore: \"rename --old-name <OLD NAME> --new-name <NEW NAME>\"",
		Action: func(c *cli.Context) {
			if err := bdevLvstoreRename(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run rename bdev lvstore command")
			}
		},
	}
}

func bdevLvstoreRename(c *cli.Context) error {
	spdkCli, err := client.NewClient()
	if err != nil {
		return err
	}

	renamed, err := spdkCli.BdevLvolRenameLvstore(c.String("old-name"), c.String("new-name"))
	if err != nil {
		return err
	}

	return util.PrintObject(renamed)
}

func BdevLvstoreDeleteCmd() cli.Command {
	return cli.Command{
		Name: "delete",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "lvs-name",
				Usage: "Specify this or uuid",
			},
			cli.StringFlag{
				Name:  "uuid",
				Usage: "Specify this or lvs-name",
			},
		},
		Usage: "delete a bdev lvstore using a block device: \"delete --lvs-name <LVSTORE NAME>\" or \"delete --uuid <UUID>\"",
		Action: func(c *cli.Context) {
			if err := bdevLvstoreDelete(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run delete bdev lvstore command")
			}
		},
	}
}

func bdevLvstoreDelete(c *cli.Context) error {
	spdkCli, err := client.NewClient()
	if err != nil {
		return err
	}

	deleted, err := spdkCli.BdevLvolDeleteLvstore(c.String("lvs-name"), c.String("uuid"))
	if err != nil {
		return err
	}

	return util.PrintObject(deleted)
}

func BdevLvstoreGetCmd() cli.Command {
	return cli.Command{
		Name: "get",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "lvs-name",
				Usage: "If you want to get one specific Lvstore info, please input this or uuid",
			},
			cli.StringFlag{
				Name:  "uuid",
				Usage: "If you want to get one specific Lvstore info, please input this or lvs-name",
			},
		},
		Usage: "get all bdev lvstore if the info is not specified: \"get\", or \"get --lvs-name <LVSTORE NAME>\", or \"get --uuid <UUID>\"",
		Action: func(c *cli.Context) {
			if err := bdevLvstoreGet(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run get bdev lvstore command")
			}
		},
	}
}

func bdevLvstoreGet(c *cli.Context) error {
	spdkCli, err := client.NewClient()
	if err != nil {
		return err
	}

	bdevLvstoreGetResp, err := spdkCli.BdevLvolGetLvstore(c.String("lvs-name"), c.String("uuid"))
	if err != nil {
		return err
	}

	return util.PrintObject(bdevLvstoreGetResp)
}

func BdevLvstoreGetLvolsCmd() cli.Command {
	return cli.Command{
		Name: "list-lvols",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "lvs-name",
				Usage: "If you want to get one specific Lvstore info, please input this or uuid",
			},
			cli.StringFlag{
				Name:  "uuid",
				Usage: "If you want to get one specific Lvstore info, please input this or lvs-name",
			},
		},
		Usage: "list all logical volumes info: \"list\", or \"list --lvs-name <LVSTORE NAME>\", or \"list --uuid <LVSTORE UUID>\"",
		Action: func(c *cli.Context) {
			if err := bdevLvolList(c); err != nil {
				logrus.WithError(err).Fatalf("Failed to run list lvol command")
			}
		},
	}
}

func bdevLvolList(c *cli.Context) error {
	spdkCli, err := client.NewClient()
	if err != nil {
		return err
	}

	bdevLvstoreGetResp, err := spdkCli.BdevLvolGetLvols(c.String("lvs-name"), c.String("uuid"))
	if err != nil {
		return err
	}

	return util.PrintObject(bdevLvstoreGetResp)
}
