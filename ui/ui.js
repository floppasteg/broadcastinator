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

/**
 * Set the top banner
 * @param {string} status - status to set 
 * @param {string} [extra] - optional extra text
 */
function changeStatus(status, extra) {
    let elem = document.getElementById("status")
    elem.classList.remove("blink-tobg", "breathe-tobg", "blink-tobgfast")

    switch (status) {
        case "disconnected":
            elem.style.backgroundColor = "orange"
            elem.innerText = "Disconnected"
            break
        case "connected-inactive":
            elem.style.backgroundColor = "yellowgreen"
            elem.style.color = "black"
            elem.innerText = "Connected"
            break;
        case "connected-live":
            elem.style.backgroundColor = "red"
            elem.style.color = "black"
            elem.innerText = "Live"
            elem.classList.add("breathe-tobg")
            break;
        case "connected-servererror":
            elem.style.backgroundColor = "red"
            elem.style.color = "black"
            elem.innerText = "Server error"
            elem.classList.add("blink-tobg")
            break;
        case "sockerror":
            elem.style.backgroundColor = "red"
            elem.style.color = "black"
            elem.innerText = "Socket error"
            elem.classList.add("blink-tobgfast")
            break;
        default:
            console.error("unknown status %s", status)
    }
    if (extra != "")
        elem.innerText += " (" + extra + ")"
}

/**
 * Add a line to the log display
 * @param {string} logline 
 * @param {string} severity 
 */
function updateLog(logline, severity) {
    let escapedLogline = new Option(logline).innerHTML
    let timestamp = new Date()
    let line = `<span style="color:gray">[${timestamp.getFullYear()}-${timestamp.getMonth()}-${timestamp.getDate()} ${timestamp.getHours()}:${timestamp.getMinutes()}:${timestamp.getSeconds()}.${timestamp.getMilliseconds()}]</span> `
    switch (severity) {
        case "debug":
            line += `<span style="color:#5a58f4">${escapedLogline}</span>\n`
            break;
        case "info":
            line += `<span>${escapedLogline}</span>\n`
            break;
        case "warn":
            line += `<span style="color:yellow">${escapedLogline}</span>\n`
            break;
        case "error":
            line += `<span style="color:red">${escapedLogline}</span>\n`
            break;
    }
    let elem = document.getElementById('log1-container')
    elem.innerHTML+=line
    if(followingLog) {
        elem.scrollTo(0,elem.scrollHeight)
    }
    /**@type {StoredLogline[]} */
    let prevlines = JSON.parse(localStorage.getItem("log1-prev"))
    if (prevlines != null) {
        prevlines.push({ timestamp: timestamp.getTime(), message: logline, severity: severity })
        localStorage.setItem('log1-prev', JSON.stringify(prevlines))
    } else {
        prevlines = [{ timestamp: timestamp.getTime(), message: logline, severity: severity }]
        localStorage.setItem('log1-prev', JSON.stringify(prevlines))
    }
}

function restoreLog() {
    /**@type {StoredLogline[]} */
    let prevlines = JSON.parse(localStorage.getItem("log1-prev"))
    let elem = document.getElementById('log1-container')
    for (let i = 0; i < prevlines.length; i++) {
        let tn = new Date(prevlines[i].timestamp).toISOString()
        let escapedLogline = new Option(prevlines[i].message).innerHTML
        elem.innerHTML += `<span style="color:gray">[${tn}]</span> ${escapedLogline}\n`
    }
}

function clearLog() {
    localStorage.removeItem('log1-prev')
    localStorage.setItem('log1-prev',"[]")
    document.getElementById('log1-container').innerHTML = ""
}

var followingLog=true;
function followLog() {
    followingLog=!followingLog
    if(followingLog) {
        document.getElementById('log1-follow').style.backgroundColor='darkgray'
        document.getElementById('log1-follow').style.color='white'
    } else {
        document.getElementById('log1-follow').style.backgroundColor=''
        document.getElementById('log1-follow').style.color=''
    }
}

function environ() {
    let noWebsocket = new URL(location.href).searchParams.get("nows")
    if (location.hostname == "localhost" || location.hostname == "127.0.0.1")
        return "online"
    else if (noWebsocket) return "online-nows"
    else return "offline"
}