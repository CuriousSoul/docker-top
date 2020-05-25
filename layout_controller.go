package main

import (
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

// LayoutController hold details about screen layout
type LayoutController struct {
	drawbles []ui.Drawable
}

// NewLayoutController creates new instance
func NewLayoutController() (*LayoutController, error) {
	return &LayoutController{}, nil
}

// getInitialLayout provides the initial layout
func (cntrl *LayoutController) getInitialLayout() []ui.Drawable {
	var layout []ui.Drawable
	// width, height := ui.TerminalDimensions()
	// l := widgets.NewList()
	// l.Title = "List"
	// l.Rows = []string{
	// 	"[0] github.com/gizak/termui/v3",
	// 	"[1] [你好，世界](fg:blue)",
	// 	"[2] [こんにちは世界](fg:red)",
	// 	"[3] [color](fg:white,bg:green) output",
	// 	"[4] output.go",
	// 	"[5] random_out.go",
	// 	"[6] dashboard.go",
	// 	"[7] foo",
	// 	"[8] bar",
	// 	"[9] baz",
	// }
	// l.TextStyle = ui.NewStyle(ui.ColorYellow)
	// l.WrapText = false
	// l.SetRect(0, 0, width, height)
	// cntrl.drawbles = append(cntrl.drawbles, l)

	nodes := []*widgets.TreeNode{
		{
			Value: nodeValue("Key 1"),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue("Key 1.1"),
					Nodes: []*widgets.TreeNode{
						{
							Value: nodeValue("Key 1.1.1"),
							Nodes: nil,
						},
						{
							Value: nodeValue("Key 1.1.2"),
							Nodes: nil,
						},
					},
				},
				{
					Value: nodeValue("Key 1.2"),
					Nodes: nil,
				},
			},
		},
		{
			Value: nodeValue("Key 2"),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue("Key 2.1"),
					Nodes: nil,
				},
				{
					Value: nodeValue("Key 2.2"),
					Nodes: nil,
				},
				{
					Value: nodeValue("Key 2.3"),
					Nodes: nil,
				},
			},
		},
		{
			Value: nodeValue("Key 3"),
			Nodes: nil,
		},
	}
	l := widgets.NewTree()
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetNodes(nodes)

	x, y := ui.TerminalDimensions()

	l.SetRect(0, 0, x, y)

	cntrl.drawbles = append(cntrl.drawbles, l)
	return layout
}

// HandleEvent handles the screen events (key/mouse events)
func (cntrl *LayoutController) HandleEvent(e ui.Event) []ui.Drawable {
	switch e.Type {
	case ui.KeyboardEvent:
		switch e.ID {
		case "q":
			os.Exit(0)
		}
	case ui.MouseEvent:
		fallthrough
	case ui.ResizeEvent:
	}
	return cntrl.Refresh()
}

// Refresh provides updated layout
func (cntrl *LayoutController) Refresh() []ui.Drawable {
	if len(cntrl.drawbles) == 0 {
		return cntrl.getInitialLayout()
	}
	return cntrl.drawbles
}
