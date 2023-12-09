package utils

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func PlayVideo(link string) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	ctx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	// create context
	ctx, cancel := chromedp.NewContext(
		ctx,
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(link),
		chromedp.WaitVisible(`.ytp-fullscreen-button`),
		chromedp.Click(`.ytp-fullscreen-button`),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)

	return nil
}
