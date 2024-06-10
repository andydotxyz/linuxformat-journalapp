package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const dateFormat = "2 Jan 2006"

func main() {
	a := app.NewWithID("com.example.journal")
	w := a.NewWindow("My Journal")
	w.SetPadded(false)

	var date time.Time

	entry := widget.NewMultiLineEntry()
	title := widget.NewLabel("Today")
	title.Alignment = fyne.TextAlignCenter

	var updateMoods func()
	setMood := func(mood string) {
		dateStr := date.Format(dateFormat)
		a.Preferences().SetString(dateStr+".mood", mood)
		updateMoods()
	}

	mood := container.NewHBox(
		widget.NewButton("😀", func() {
			setMood("happy")
		}),
		widget.NewButton("😍", func() {
			setMood("loved")
		}),
		widget.NewButton("😟", func() {
			setMood("sad")
		}),
		widget.NewButton("😴", func() {
			setMood("tired")
		}),
		widget.NewButton("😬", func() {
			setMood("stressed")
		}),
	)
	updateMoods = func() {
		setSelected := func(o fyne.CanvasObject, high bool) {
			b, ok := o.(*widget.Button)
			if !ok {
				return
			}

			if high {
				b.Importance = widget.HighImportance
			} else {
				b.Importance = widget.MediumImportance
			}
			b.Refresh()
		}

		dateStr := date.Format(dateFormat)
		dayMood := a.Preferences().String(dateStr + ".mood")
		setSelected(mood.Objects[0], dayMood == "happy")
		setSelected(mood.Objects[1], dayMood == "loved")
		setSelected(mood.Objects[2], dayMood == "sad")
		setSelected(mood.Objects[3], dayMood == "tired")
		setSelected(mood.Objects[4], dayMood == "stressed")
	}

	setDate := func(d time.Time) {
		date = d
		dateStr := date.Format(dateFormat)
		title.SetText(dateStr)
		entry.Bind(binding.BindPreferenceString(dateStr, a.Preferences()))
		entry.Validator = nil

		updateMoods()
	}
	setDate(time.Now())

	prev := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		setDate(date.Add(time.Hour * -24))
	})
	next := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		setDate(date.Add(time.Hour * 24))
	})
	bar := container.NewBorder(nil, nil, prev, next, title)

	w.SetContent(container.NewBorder(bar, container.NewCenter(mood), nil, nil, entry))
	w.Resize(fyne.NewSize(240, 210))
	w.ShowAndRun()
}
