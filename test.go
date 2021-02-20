package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "SCREAM",
		Icon:    "icon.ico",
		MinSize: Size{Width: 600, Height: 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					// outTE.SetText(strings.ToUpper(inTE.Text()))
					walk.MsgBox(&walk.Dialog{}, "Hello", "World", walk.MsgBoxOK)
				},
			},
		},
	}.Run()
}
