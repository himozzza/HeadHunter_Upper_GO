package modules

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func Chrome(inputLogin, inputPassword string) {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var data string
	url := "https://hh.ru/applicant/resumes?hhtmFromLabel=header&hhtmFrom=main"
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second),
		chromedp.WaitVisible(`//*[@data-qa="expand-login-by-password"]`, chromedp.NodeVisible),
		chromedp.Click(`//*[@data-qa="expand-login-by-password"]`, chromedp.NodeVisible),

		chromedp.SendKeys(`//*[@data-qa="login-input-username"]`, inputLogin),
		chromedp.SendKeys(`//*[@data-qa="login-input-password"]`, inputPassword),
		chromedp.Click(`//*[@data-qa="account-login-submit"]`, chromedp.NodeVisible),
		chromedp.Sleep(4*time.Second),
		chromedp.OuterHTML(`html`, &data, chromedp.BySearch),
	)
	if err != nil {
		fmt.Println("Обновлено ранее.")
	}
	re := regexp.MustCompile(`data-qa="resume-title">(.*?)</span>`)
	resume := re.FindAllString(data, -1)
	for _, i := range resume {
		clone, cancel := chromedp.NewContext(ctx)
		defer cancel()
		chromedp.Run(clone,
			chromedp.Navigate(url),
			chromedp.Sleep(2*time.Second),
			chromedp.Click(`//button[normalize-space()="Поднять в поиске"]`, chromedp.NodeVisible),
			chromedp.Sleep(2*time.Second),
		)
		resumeName := strings.SplitN(strings.SplitN(i, ">", -1)[1], "<", -1)[0]
		fmt.Printf("Резюме '%s' поднято в поиске! (Логин: %s)\n", resumeName, inputLogin)
	}
}
