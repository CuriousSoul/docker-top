package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// DefaultLayout carries default structure
type DefaultLayout struct {
	docker DockerInterface
	layout *widgets.Tree
}

type cNode struct {
	id, name, state string
}

func (c cNode) String() string {
	return fmt.Sprintf("%10s: , %10s: , %10s: ,", c.id[:10], c.name, c.state)
}

func generateView(docker DockerInterface) *widgets.Tree {
	containers := docker.ListContainers()
	nodes := make([]*widgets.TreeNode, len(containers))

	for _, c := range containers {
		nodes[0] = &widgets.TreeNode{
			Value:    cNode{id: c.ID, name: c.Names[0]},
			Expanded: true,
		}
	}
	l := widgets.NewTree()
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetNodes(nodes)
	x, y := ui.TerminalDimensions()
	l.SetRect(0, 0, x, y)
	return l
}

// InitDefaultLayout initializes default layout
func InitDefaultLayout(docker DockerInterface) *DefaultLayout {
	return &DefaultLayout{
		docker: docker,
		layout: generateView(docker),
	}
}

// HandleEvent handles the keyboard/mouse event for layout
func (d *DefaultLayout) HandleEvent(e ui.Event) LayoutAction {
	switch e.Type {
	case ui.KeyboardEvent:
		switch e.ID {
		case "q":
			return ActionExit
		}
	case ui.MouseEvent:
		fallthrough
	case ui.ResizeEvent:
	}
	return ActionNoop
}

// Refresh provides updated layout
func (d *DefaultLayout) Refresh() []ui.Drawable {
	//	go generateView(d.docker)
	return []ui.Drawable{d.layout}
}
