package main

import (
	"fmt"
	"github.com/emachadomartins/go-scrap/errs"
	"github.com/playwright-community/playwright-go"
)

func main() {
	server := errs.TryCatch(
		playwright.Run(),
	)
	browser := errs.TryCatch(
		server.Chromium.Launch(),
	)
	defer func() {
		errs.ThrowOnError(
			browser.Close(),
		)
	}()

	fmt.Println("Hello World")
}
