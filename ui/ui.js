/**
 * Set the top banner
 * @param {string} status - status to set 
 * @param {string} [extra] - optional extra text
 */
function changeStatus(status, extra) {
    let elem = document.getElementById("status")
        elem.classList.remove("blink-tobg","breathe-tobg","blink-tobgfast")
  
    switch (status) {
        case "disconnected":
            elem.style.backgroundColor = "orange"
            elem.innerText = "Disconnected (" + extra + ")"
            break
        case "connected-inactive":
            elem.style.backgroundColor = "yellowgreen"
            elem.style.color = "black"
            elem.innerText = "Connected (" + extra + ")"
            break;
        case "connected-live":
            elem.style.backgroundColor = "red"
            elem.style.color = "black"
            elem.innerText = "Live (" + extra + ")"
            elem.classList.add("breathe-tobg")
            break;
        case "connected-servererror":
            elem.style.backgroundColor = "red"
            elem.style.color = "black"
            elem.innerText = "Server error: " + extra
            elem.classList.add("blink-tobg")
            break;
        case "sockerror":
            elem.style.backgroundColor = "red"
            elem.style.color = "black"
            elem.innerText = "Socket error: " + extra
            elem.classList.add("blink-tobgfast")
            break;
        default:
            console.error("unknown status %s", status)
    }
}