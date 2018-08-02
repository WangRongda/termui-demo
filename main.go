package main

import (
	"fmt"

	ui "github.com/gizak/termui"
) // <- ui shortcut, optional

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	strs := []string{
		"[0]Something went wrong",
		"[1] editbox.go",
		"[2] interrupt.go",
		"[3] keyboard.go",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go",
		"[8] editbox.go",
		"[9] interrupt.go",
		"[10] keyboard.go",
		"[11] output.go",
		"[12] random_out.go",
		"[13] dashboard.go",
		"[14] nsf/termbox-go",
		"[15] editbox.go",
		"[16] interrupt.go",
		"[17] keyboard.go",
		"[18] output.go",
		"[19] random_out.go",
		"[20] dashboard.go",
	}
	l := ui.NewList()
	l.Items = strs
	l.ItemFgColor = ui.ColorBlack
	// l.BorderLabel = "List"
	l.Y = 0
	l.Height = len(strs)
	l.Width = 30
	l.Border = false
	l.ItemBgColor = ui.ColorYellow

	n := 0 //当前行, selected num
	shift := 0
	oldLine := l.Items[0]
	selectedWrap := "[%s](bg-red,fg-bold)"
	l.Items[0] = fmt.Sprintf(selectedWrap, l.Items[0])
	ui.Render(l) // feel free to call Render, it's async and non-block

	// event handler...
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/j", func(ui.Event) {
		if n >= shift+ui.TermHeight()-2 && shift+ui.TermHeight() < len(strs) {
			shift++
			l.Y = -shift
		} else if n >= len(l.Items)-1 {
			return
		}
		l.Items[n] = oldLine
		n++
		oldLine = l.Items[n]
		l.Items[n] = fmt.Sprintf(selectedWrap, l.Items[n])
		ui.Render(l)
	})
	ui.Handle("/sys/kbd/k", func(ui.Event) {
		if n == shift+1 && shift > 0 {
			shift--
			l.Y = -shift
		} else if n <= 0 {
			return
		}
		l.Items[n] = oldLine
		n--
		oldLine = l.Items[n]
		l.Items[n] = fmt.Sprintf(selectedWrap, l.Items[n])
		ui.Render(l)
	})
	ui.Handle("/sys/wnd/resize", func(ui.Event) {
		// ui.Body.Height = ui.TermHeight()
		// ui.Body.Align()
		ui.Clear()
		ui.Render(l)
	})
	ui.Loop() // block until StopLoop is called
}
