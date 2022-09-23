package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/chromedp/chromedp"
)

func main() {
	var data string
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("headless", false), // debug
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx) // enable debug log to see the CDP traffics.
	//chromedp.WithDebugf(log.Printf),

	defer cancel()

	url := "https://hh.ru/applicant/resumes?hhtmFromLabel=header&hhtmFrom=main"

	chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Click(`//*[@data-qa="expand-login-by-password"]`, chromedp.NodeVisible),
		// chromedp.Sleep(15*time.Second),
		chromedp.SendKeys(`//*[@data-qa="login-input-username"]`, login),
		chromedp.SendKeys(`//*[@data-qa="login-input-password"]`, passwd),
		chromedp.Click(`//*[@data-qa="account-login-submit"]`, chromedp.NodeVisible),
		chromedp.WaitVisible(`//button[@data-qa="resume-update-button_actions"]`, chromedp.NodeVisible),
		// chromedp.Click(`//*[]`),
		chromedp.OuterHTML("html", &data, chromedp.ByQuery),
		// chromedp.Sleep(10*time.Second),
	)
	a := reg(data)
	fmt.Println(a)
}
func reg(data string) string {
	re := regexp.MustCompile(`data-qa="resume-update-button_actions"`)
	a := re.FindString(data)
	return a
}
