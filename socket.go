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
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Socket struct {
	sock *websocket.Conn
}

func initWebsocket() *Socket {
	sock := new(Socket)
	upg := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	http.HandleFunc("/crunchy", func(w http.ResponseWriter, r *http.Request) {
		con, err := upg.Upgrade(w, r, nil)
		chk(err)
		sock.sock = con
	})

	return sock
}

func (s *Socket) printf(format string, v ...any) {
	w, e := s.sock.NextWriter(websocket.TextMessage)
	chk(e)
	fmt.Fprintf(w, format, v...)
}

// For now we use the websocket to send messages to the web interface
