package tapper

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Application struct {
	screen tcell.Screen

	gState        *_globalstate
	signalChannel chan Signal
}

func (app *Application) Run() {
	app.screen.Init()
	app.screen.EnableFocus()
	go app._backgroundEventHandler()
	app.mainLoop()
}

func (app *Application) Stop() {
	app.signalChannel <- Signal{
		Sigtype: SignalQuit}
}

func (app *Application) TestSendCallback() {
	app.signalChannel <- Signal{
		Sigtype: SignalCallback,
		Data: func() {
			app.screen.SetContent(0, 0, tcell.RuneBlock, nil, tcell.StyleDefault.Blink(true))
		},
	}
}

func (app *Application) _backgroundEventHandler() {
	for {
		event := app.screen.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				app.Stop()
				return
			case tcell.KeyEnter:
				app.TestSendCallback()
			}
		case *tcell.EventResize:
			app.screen.Sync()
		case *tcell.EventFocus:
			app.signalChannel <- Signal{
				Sigtype: SignalFocus,
				Data:    ev.Focused,
			}

		}
	}
}
func _mainloopEventHandler(screen tcell.Screen, sigChan chan Signal) {
	receivedSignal := <-sigChan
	switch receivedSignal.Sigtype {
	case SignalQuit:
		screen.Fini()
		os.Exit(0)
	case SignalCallback:
		// fmt.Println(receivedSignal.Data)
		receivedSignal.Data.(func())()
		// app.screen.SetContent(0, 0, 'a', nil, tcell.StyleDefault)
		// callback(app.screen)
	case SignalFocus:
		displayString := 'f'
		if receivedSignal.Data.(bool) {
			displayString = 'c'
		}
		screen.SetContent(1, 1, displayString, nil, tcell.StyleDefault.Foreground(tcell.Color101).Background(tcell.ColorBlack))
	case SignalDraw:
		screen.Sync()
	}
}
func (app *Application) mainLoop() {
	box := NewBox(0, 0, 10, 10, "Hello World!")
	box.SetFocus(true)
	app.screen.Clear()
	for {
		// app.screen.Clear()
		box.Draw(app.screen)
		_mainloopEventHandler(app.screen, app.signalChannel)
		app.screen.Show()
	}
}

func NewApplication() *Application {

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	gState := initGlobalState()
	return &Application{
		screen: s,
		gState: gState,

		signalChannel: make(chan Signal),
	}
}
