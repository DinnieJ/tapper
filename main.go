//go:build ignore
// +build ignore

// Copyright 2022 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// mouse displays a text box and tests mouse interaction.  As you click
// and drag, boxes are displayed on screen.  Other events are reported in
// the box.  Press ESC twice to exit the program.
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/DinnieJ/tapper"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	// "github.com/mattn/go-runewidth"
)

func main() {

	shell := os.Getenv("SHELL")
	if shell == "" {
		if runtime.GOOS == "windows" {
			shell = "CMD.EXE"
		} else {
			shell = "/bin/sh"
		}
	}

	encoding.Register()

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	// defStyle = tcell.StyleDefault.
	// 	Background(tcell.ColorReset).
	// 	Foreground(tcell.ColorReset)
	// s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.EnableFocus()
	// s.Clear()
	box := tapper.NewBox(5, 5, 10, 10, "fucked")
	quit := make(chan int)
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {

				case tcell.KeyEscape, tcell.KeyEnter:
					quit <- 0
					return
				case tcell.KeyCtrlB:
					quit <- 1
				case tcell.KeyRune:

				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()
	for {
		s.Clear()
		select {
		case value := <-quit:
			if value == 1 {
				box.Y1 += 1
			} else {
				s.Fini()
				os.Exit(0)
				break
			}
		case <-time.After(time.Millisecond * 100):
		}
		box.X1 += 1
		// box.Y1 += 1
		box.X2 += 1

		box.Draw(s)
		// ev := s.PollEvent()
		// switch ev.(type) {
		// case *tcell.EventResize:
		// 	s.Sync()
		// }

		s.Show()
		// time.Sleep(300 * time.Millisecond)
	}

}
