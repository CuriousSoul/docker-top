package main

import (
	"os"

	ui "github.com/gizak/termui/v3"
)

// LayoutAction indicates operations related to
type LayoutAction int

const (
	// ActionNoop action -- no change
	ActionNoop LayoutAction = iota + 1
	// ActionExit action -- Layout exit
	ActionExit
)

// Layout interface definition
type Layout interface {
	HandleEvent(e ui.Event) LayoutAction
	Refresh() []ui.Drawable
}

// LayoutController hold details about screen layout
type LayoutController struct {
	layouts []Layout
}

// NewLayoutController creates new instance
func NewLayoutController(docker DockerInterface) (*LayoutController, error) {
	cntrl := LayoutController{
		layouts: []Layout{InitDefaultLayout(docker)},
	}
	return &cntrl, nil
}

// HandleEvent handles the screen events (key/mouse events)
func (cntrl *LayoutController) HandleEvent(e ui.Event) []ui.Drawable {
	switch action := cntrl.layouts[0].HandleEvent(e); action {
	case ActionExit:
		os.Exit(0)
	case ActionNoop:
	}
	return cntrl.layouts[0].Refresh()
}

// Refresh provides updated layout
func (cntrl *LayoutController) Refresh() []ui.Drawable {
	return cntrl.layouts[0].Refresh()
}
