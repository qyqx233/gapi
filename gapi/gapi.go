package main

import (
	"context"
	"flag"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ExecuteCmdRq struct {
	Cmd     string `json:"cmd"`
	Timeout int    `json:"timeout"`
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

	app := fiber.New()
	app.Post("/executeCmd", func(c *fiber.Ctx) error {
		// Parse request body
		var rq ExecuteCmdRq
		if err := c.BodyParser(&rq); err != nil {
			return err
		}
		if rq.Timeout == 0 {
			rq.Timeout = 30
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rq.Timeout)*time.Second)
		defer cancel()
		outStr, errStr, err := executeCmd(ctx, rq.Cmd)
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
