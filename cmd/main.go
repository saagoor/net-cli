package main

import (
	"log"
	"os"

	"github.com/saagoor/net-cli/cmd/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Net Lookup CLI"
	app.Usage = "Lets you query IP's, CNAME's, MX Records & Nameservers!"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "tutorialedge.net",
		},
	}

	app.Commands = []cli.Command{
		commands.NS(myFlags),
		commands.IP(myFlags),
		commands.CNAME(myFlags),
		commands.MX(myFlags),
	}

	err := app.Run((os.Args))
	if err != nil {
		log.Fatal(err)
	}

}
