package main

import (
	"context"
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
	app := fiber.New()
	app.Post("/executeCmd", func(c *fiber.Ctx) error {
		// Parse request body
		var req ExecuteCmdRq
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(req.Timeout)*time.Second)
		defer cancel()
		outStr, errStr, err := executeCmd(ctx, req.Cmd)
		// Process request

		res := ExecuteCmdRs{
			Err: errStr,
			Out: outStr,
		}

		if err != nil {
			res.Msg = err.Error()
		}
		// Return response
		return c.JSON(res)
	})

	app.Listen(":3000")
}
