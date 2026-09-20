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

const socketURL = "/ctrlsock"

/** @type {WebSocket} */
let socket;
if (environ() == "online") {
    socket = new WebSocket(socketURL)
    socket.onerror = function (event) {
        changeStatus("sockerror", "onerror triggered, check F12 log for details")
        console.error("onerror in socket.js triggered: ", event)
    }
    socket.onmessage = function(ev) {
        /** @type {SocketMessage} */
        let obj = JSON.parse(ev.data)
        switch(obj.msgtype) {
            case "log":
                updateLog(obj.msgdata.substring(4),obj.msgdata.substring(0,4))
            case "status":
                switch(obj.msgdata.substring(0,4)) {
                    case "idle":
                        changeStatus("connected-inactive","")
                        break;
                    case "live":
                        changeStatus("connected-live")
                        break;
                    case "serr":
                        changeStatus("connected-servererror",obj.msgdata.substring(4))
                }
                break;
        }
    }
}