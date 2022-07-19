package commands

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func NS(flags []cli.Flag)(cli.Command)  {
	return cli.Command{
		Name:  "ns",
		Usage: "Looks up the nameservers of a particular host.",
		Flags: flags,
		Action: func(c *cli.Context) error {
			ns, err := net.LookupNS(c.String("host"))
			if err != nil {
				fmt.Println(err)
				return err
			}
			for i := 0; i < len(ns); i++ {
				fmt.Println(ns[i].Host)
			}
			return nil
		},
	}
}