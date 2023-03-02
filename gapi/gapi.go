package main

import (
	"context"
	"flag"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ExecuteCmdRq struct {
	Timeout int    `json:"timeout"`
	Cmd     string `json:"cmd"`
	Path    string `json:"path"`
}

type ExecuteCmdRs struct {
	Msg string `json:"msg,omitempty"`
	Out string `json:"out,omitempty"`
	Err string `json:"err,omitempty"`
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "c", "gapi.yml", "config path")
	flag.Parse()
	config := parseYaml(configPath)
	if config.App.MaxCpu == 0 {
		config.App.MaxCpu = 4
	}
	runtime.GOMAXPROCS(config.App.MaxCpu)

	app := fiber.New()
	app.Post("/executeCmd", func(c *fiber.Ctx) error {
		var rq ExecuteCmdRq
		if err := c.BodyParser(&rq); err != nil {
			return err
		}
		if rq.Timeout == 0 {
			rq.Timeout = 30
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rq.Timeout)*time.Second)
		defer cancel()
		outStr, errStr, err := executeCmd(ctx, rq.Cmd, rq.Path)
		// Process request

		rs := ExecuteCmdRs{
			Err: errStr,
			Out: outStr,
		}

		if err != nil {
			rs.Msg = err.Error()
		}
		// Return response
		return c.JSON(rs)
	})

	app.Listen(config.App.Addr)
}
