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