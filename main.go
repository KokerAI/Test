package main

import (
	"fmt"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GO Калькулятор")
	w.Resize(fyne.NewSize(400, 300))

	icon, _ := fyne.LoadResourceFromPath("app.png")
	w.SetIcon(icon)

	url, _ := url.Parse("https://sun9-42.userapi.com/impf/c830309/v830309737/153902/kVXyXLeUyR4.jpg?size=762x670&quality=96&sign=b0e291b82e0ee77014e67bc29cbf3298&type=album")

	link := widget.NewHyperlink("Автор программы", url)

	label1 := widget.NewLabel("Введите первое число:")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Введите второе число:")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("")

	btn := widget.NewButton("ПОСЧИТАТЬ", func() {
		n1, err := strconv.ParseFloat(entry1.Text, 64)
		n2, err2 := strconv.ParseFloat(entry2.Text, 64)

		if err != nil || err2 != nil {
			answer.SetText("УЧИ МАТЕМАТИКУ!")
		} else {
			sum := n1 + n2
			sub := n1 - n2
			mul := n1 * n2
			div := n1 / n2

			answer.SetText(fmt.Sprintf("(+) %f\n(-) %f\n(*) %f\n(/) %f", sum, sub, mul, div))
		}
	})

	w.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		btn,
		answer,
		link,
	))

	w.ShowAndRun()
}
