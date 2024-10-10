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
func _mainloopEventHandler(app *Application, sigChan chan Signal) {
	receivedSignal := <-sigChan
	switch receivedSignal.Sigtype {
	case SignalQuit:
		app.screen.Fini()
		os.Exit(0)
	case SignalCallback:
		// fmt.Println(receivedSignal.Data)
		receivedSignal.Data.(func())()
		// app.screen.SetContent(0, 0, 'a', nil, tcell.StyleDefault)
		// callback(app.screen)
	case SignalFocus:
		boxState := app.gState.Get("test").(*Box)
		displayString := 'f'
		boxState.SetFocus(false)
		if receivedSignal.Data.(bool) {
			displayString = 'c'
			boxState.SetFocus(true)
		}

		app.screen.SetContent(1, 1, displayString, nil, tcell.StyleDefault.Foreground(tcell.Color101).Background(tcell.ColorBlack))

	case SignalDraw:
		app.screen.Sync()
	}
}
func (app *Application) mainLoop() {
	box := NewBox("test", 0, 0, 10, 10)
	app.gState.Set("test", box)
	box.SetFocus(true)
	app.screen.Clear()
	for {
		// app.screen.Clear()
		_mainloopEventHandler(app, app.signalChannel)
		box.Draw(app.screen)
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
	s.Init()
	s.EnableFocus()
	return &Application{
		screen: s,
		gState: gState,

		signalChannel: make(chan Signal),
	}
}
