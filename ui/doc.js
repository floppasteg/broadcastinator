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
 * This file's entire purpose is to add object types for autocomplete and documentation
 * purposes
 * i lovingly hate you javascript (and loosely typed languages in general)
 */

/**
 * @typedef {Object} SocketMessage A message coming from the server
 * @property {string} msgtype
 * @property {string} msgdata
 * @property {number} timestamp
 */

/**
 * @typedef {Object} StoredLogline An object containing a specific log message, complete with timestamp and 
 * @property {string} severity
 * @property {string} message
 * @property {number} timestamp
 */