package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"main.go/modules"
)

func updateTime(clock *widget.Label, i int) {
	if i == 0 {
		clock.SetText("Работаем!")
		modules.Chrome()
	} else {
		formatted := "До старта осталось: " + strconv.Itoa(i) + " секунд!"
		clock.SetText(formatted)
	}
}
func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("hh Updater GO")

	inputLogin := widget.NewEntry()
	inputPassword := widget.NewEntry()
	inputLogin.SetPlaceHolder("Логин:")
	inputPassword.SetPlaceHolder("Пароль:")
	if desk, ok := myApp.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				myWindow.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	content := container.NewVBox(inputLogin, inputPassword, widget.NewButton("Сохранить и запустить", func() {
		modules.WriteFile(inputLogin.Text, inputPassword.Text)
		clock := widget.NewLabel("")
		i := 14500
		updateTime(clock, i)
		myWindow.SetContent(clock)

		go func() {
			for range time.Tick(time.Second) {
				if i <= 0 {
					i = 14500
				} else {
					i--
					updateTime(clock, i)
				}
			}

		}()

	}))
	myWindow.SetContent(widget.NewLabel("Fyne System Tray"))
	myWindow.SetCloseIntercept(func() {
		myWindow.Hide()
	})
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
