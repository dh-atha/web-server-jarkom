package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	folder := "logs"
	filename := "server.log"
	path := fmt.Sprintf("%s/%s", folder, filename)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, 0755)
	}

	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(logFile, os.Stdout)

	logger := slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		AddSource: true,
	}))

	serverID := fmt.Sprintf("server-%d", os.Getpid())

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		msg := fmt.Sprintf("[Server ID: %s] %s %s - %d", serverID,
			c.Method(),
			c.OriginalURL(),
			c.Response().StatusCode(),
		)
		logger.Info(msg, slog.String("duration", time.Since(start).String()))
		return err
	})

	app.Static("/public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		tmpl, err := template.ParseFiles("public/index.html")
		if err != nil {
			return err
		}

		tmpl.Execute(c, map[string]interface{}{
			"ServerID": serverID,
		})

		c.Set("Content-Type", "text/html")
		return nil
	})

	logger.Info(fmt.Sprintf("Server (Server ID: %s) running on port %s", serverID, port))
	app.Listen(fmt.Sprintf(":%s", port))
}
