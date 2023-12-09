package utils

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func PlayVideo(link string) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("mute-audio", false),
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
		chromedp.Sleep(10*time.Second), // Add a delay here
		chromedp.WaitVisible(`.ytp-fullscreen-button`),
		chromedp.Click(`.ytp-fullscreen-button`),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)

	return nil
}
