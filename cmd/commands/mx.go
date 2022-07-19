package commands

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func MX(flags []cli.Flag) cli.Command {
	return cli.Command{
		Name:  "mx",
		Usage: "Looks up the MX records of a particular host.",
		Flags: flags,
		Action: func(c *cli.Context) error {
			mx, err := net.LookupMX(c.String("host"))
			if err != nil {
				println(err)
				return err
			}
			for i := 0; i < len(mx); i++ {
				fmt.Println(mx[i].Host, mx[i].Pref)
			}
			return nil
		},
	}
}
