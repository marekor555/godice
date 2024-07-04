package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	SIDES = 6
)

var (
	results     = map[int]int{}
	scoreBoard  *widget.RichText
	result      *canvas.Text
	rollAmount  *widget.Select
	rollAmounts = []string{
		"3",
		"5",
		"10",
		"12",
		"20",
		"50",
		"75",
		"100",
		"1000",
		"10000",
		"100000",
		"1000000",
	}
)

func refreshScoreBoard() {
	segments := []widget.RichTextSegment{}
	for i := 1; i <= SIDES; i++ {
		segments = append(segments, &widget.TextSegment{
			Text: fmt.Sprintf("%v - %v", i, results[i]),
		})
	}
	scoreBoard.Segments = segments
	scoreBoard.Refresh()
}

func resetScoreBoard() {
	for i := 1; i <= SIDES; i++ {
		results[i] = 0
	}
	refreshScoreBoard()
}

func rollDice(animation bool) {
	if animation {
		for range 10 {
			result.Text = fmt.Sprint(rand.Intn(SIDES) + 1)
			result.Refresh()
			time.Sleep(time.Millisecond * 20)
		}
	}
	finalResult := rand.Intn(SIDES) + 1
	if animation {
		result.Text = fmt.Sprint(finalResult)
		result.Refresh()
		refreshScoreBoard()
	}
	results[finalResult]++
}

func main() {
	os.Setenv("FYNE_THEME", "dark")
	application := app.New()
	window := application.NewWindow("godice - dice simulator")
	window.Resize(fyne.NewSize(350, 350))

	result = canvas.NewText("", theme.ForegroundColor())
	result.TextStyle = fyne.TextStyle{Bold: true}
	result.TextSize = 50
	result.Refresh()

	scoreBoard = widget.NewRichText()
	scoreBoard.Refresh()

	resetScoreBtn := widget.NewButton("reset", resetScoreBoard)

	resultAlign := container.NewHBox(layout.NewSpacer(), result, layout.NewSpacer())

	rollOnce := widget.NewButton("roll the dice", func() {
		go rollDice(true)
	})

	rollMany := widget.NewButton("roll the amount", func() {
		amount, err := strconv.Atoi(rollAmount.Selected)
		if err != nil {
			result.Text = err.Error()
			result.Refresh()
			return
		}
		result.Text = "Calculating..."
		result.Refresh()

		for i := 0; i < amount; i++ {
			rollDice(false)
		}
		refreshScoreBoard()

		result.Text = "Done"
		result.Refresh()

	})

	rollAmount = widget.NewSelect(rollAmounts, func(s string) {})
	rollAmount.SetSelected(rollAmounts[0])

	rollManyCont := container.NewGridWithColumns(2, rollAmount, rollMany)

	vbox := container.NewVBox(scoreBoard, resetScoreBtn)

	mainCont := container.NewVBox(rollManyCont, rollOnce, layout.NewSpacer(), resultAlign, layout.NewSpacer(), vbox)

	refreshScoreBoard()
	for i := 1; i <= SIDES; i++ {
		results[i] = 0
	}

	window.SetContent(mainCont)
	window.ShowAndRun()
}
