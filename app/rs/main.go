package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zhiqiangxu/rustscript/pkg/vm"
)

func main() {

	app := &cli.App{
		Name:  "rs",
		Usage: "rs wasm_file",
		Action: func(c *cli.Context) (err error) {
			if c.NArg() != 1 {
				err = fmt.Errorf("wasm_file missing")
				return
			}
			wsFile := c.Args().Get(0)
			code, err := ioutil.ReadFile(wsFile)
			if err != nil {
				return
			}

			engine := vm.NewEngine(code)
			ret, err := engine.Execute()
			fmt.Println(ret, err)
			return
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
