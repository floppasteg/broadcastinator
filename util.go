/*
   broadcastinator: an RTMP TV channel simulator
   Copyright (C) 2026  floppasteg

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

type globalStateStruct struct {
	logger *slog.Logger
}

var (
	globalLogger *slog.Logger
	globalState  globalStateStruct
)

func chk(e error, panicable bool) {
	if e != nil {
		globalLogger.Error(fmt.Sprintf("error caught in chk(e error): %v", e))
		if panicable {
			panic(e)
		}
	}
}

func openOrCreate(path string) *os.File {
	f, e := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0777)
	if errors.Is(e, os.ErrNotExist) {
		f, e = os.Create(path)
		chk(e, true)
		return f
	} else {
		chk(e, true)
	}
	println(f == nil)
	// check for "bad file descriptor" issues before passing it on to Gin and slog where the error
	// won't pop up and cause the logfile to not write new logs
	// fixed with changing os.Open to os.OpenFile (should probably also do it to autocopy's checkpoint
	// file logic)
	_, e = fmt.Fprintln(f)
	if e != nil {
		panic(e)
	}
	return f
}

func assert(cond bool, format string, v ...any) {
	if !cond {
		panic(fmt.Errorf(format, v...))
	}
}
