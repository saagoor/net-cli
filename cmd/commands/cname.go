package commands

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func CNAME(flags []cli.Flag) (cli.Command) {
	return cli.Command{
		Name:  "cname",
		Usage: "Looks up the cname records of a particular host.",
		Flags: flags,
		Action: func(c *cli.Context) error {
			cname, err := net.LookupCNAME(c.String("host"))
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(cname)
			return nil
		},
	}
}