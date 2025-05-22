package main

import (
	"github.com/emachadomartins/go-scrap/errs"
	"github.com/playwright-community/playwright-go"
	"os"
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

	page := errs.TryCatch(
		browser.NewPage(),
	)

	defer func() {
		errs.ThrowOnError(
			page.Close(),
		)
	}()

	errs.TryCatch(
		page.Goto("https://example.com"),
	)

	html := errs.TryCatch(
		page.Content(),
	)

	errs.ThrowOnError(
		os.MkdirAll("./output", 0755),
	)

	errs.ThrowOnError(
		os.WriteFile(
			"./output/example.html",
			[]byte(html),
			0644),
	)
}
