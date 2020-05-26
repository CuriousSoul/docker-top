package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	docker, err := NewDockerClient()
	if err != nil {
		log.Fatalf("Could not initialize client")
	}

	layoutCntrl, err := NewLayoutController(docker)
	if err != nil {
		log.Fatalf("Could not initialize client")
	}

	ticker := time.Tick(10 * time.Millisecond)
	for {
		select {
		case e := <-ui.PollEvents():
			ui.Render(layoutCntrl.HandleEvent(e)...)
		case <-ticker:
			ui.Render(layoutCntrl.Refresh()...)
		}
	}
}
