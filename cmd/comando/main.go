package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/urfave/cli"

	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Value: "echo 'as you were!'",
			Usage: "Command to execute",
		},
		cli.StringFlag{
			Name:  "r",
			Value: "comando",
			Usage: "Route to serve on the command output",
		},
		cli.BoolFlag{
			Name: "raw-output",
			Usage: `Output will be returned directly as a JSON string
	rather than being trimmed.`,
		},
		cli.StringFlag{
			Name:  "port",
			Value: "6000",
			Usage: "Set port",
		},
		cli.StringFlag{
			Name:  "host",
			Value: "localhost",
			Usage: "Set host",
		},
	}

	app.Name = "comando"
	app.Usage = "Command as a service"
	app.Action = func(c *cli.Context) error {
		var err error
		r := mux.NewRouter()
		r.HandleFunc("/"+c.String("r"), func(w http.ResponseWriter, r *http.Request) {
			out, err := exec.Command("bash", "-c", c.String("c")).Output()
			if err != nil {
				return
			}

			data := map[string]string{}
			if c.Bool("raw-output") {
				data["output"] = string(out)
			} else {
				data["output"] = strings.TrimSpace(string(out))
			}
			json.NewEncoder(w).Encode(data)
			return
		}).Methods("GET")

		if err != nil {
			return err
		}
		addr := fmt.Sprintf("%s:%s", c.String("host"), c.String("port"))
		http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, r))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("run app: %s", err)
	}
}
