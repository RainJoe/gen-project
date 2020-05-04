package main

import (
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "gen-project",
		Usage: "Go web 新项目创建工具",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "d",
				Value:       "",
				Usage:       "指定项目所在目录",
				Destination: &p.path,
			},
			&cli.StringFlag{
				Name:        "m",
				Value:       "",
				Usage:       "go module目录",
				Destination: &p.ModPrefix,
			},
		},
	}
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		if err := app.Run([]string{"-h"}); err != nil {
			log.Fatal(err)
		}
		return
	}
	app.Action = runNew
	p.Name = os.Args[1]
	args := append([]string{os.Args[0]}, os.Args[2:]...)
	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
