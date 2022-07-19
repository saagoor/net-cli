package commands

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func IP(flags []cli.Flag)(cli.Command)  {
	return cli.Command{
		Name:  "ip",
		Usage: "Looks up the ip addresses of a particular host.",
		Flags: flags,
		Action: func(c *cli.Context) error {
			ip, err := net.LookupIP(c.String("host"))
			if err != nil {
				fmt.Println(err)
				return err
			}
			for i := 0; i < len(ip); i++ {
				fmt.Println(ip[i])
			}
			return nil
		},
	}
}