package main

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"time"

	"github.com/labstack/echo"
)

var source = rand.NewSource(time.Now().UnixNano())
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

func main() {
	generatedStr := RandString(1000000)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		c.Response().WriteHeader(http.StatusOK)

		enc := json.NewEncoder(c.Response())
		for {
			if err := enc.Encode(generatedStr); err != nil {
				return err
			}
			c.Response().Flush()
			//time.Sleep(1 * time.Millisecond)
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":3009"))
}
