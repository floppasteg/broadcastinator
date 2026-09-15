const masterCtlAPIPath = "/masterctrl"

/**
 * Send a master command
 * @param {string} cmd - either "start", "override", "shutdown-soft", "shutdown-hard", "e-stop"
 */
function sendMasterCommand(cmd) {
    console.log(location.hostname)
    if (location.hostname != 'localhost' && location.hostname != '127.0.0.1') {
        console.log("assuming debugging in operation... setting status instead");
        masterCommandToStatus(cmd)
        return
    }
    console.log("sending master command %s...", cmd)
    let req = new XMLHttpRequest()
    req.open("POST", "http://" + location.hostname+ ":8088" + masterCtlAPIPath + "/" + cmd)
    req.addEventListener('error', function () {
        console.error("could not send master command")
    })
    req.addEventListener('load', function () {
        let obj = JSON.parse(this.responseText)
        if (obj.status != 'success') {
            console.error("MasterCommand.status != success, details: ", obj)
        }
        masterCommandToStatus(cmd)
    })
    req.addEventListener('error', function () {
        console.error("")
    })
    req.send()
}

function masterCommandToStatus(cmd, extra) {
    switch (cmd) {
        case 'start':
            changeStatus('connected-live', "now playing...")
            break
        case 'e-stop':
            changeStatus('disconnected', "E-stop hit")
            break
        case 'shutdown-soft':
            changeStatus('connected-inactive', "soft shutdown triggered")
            break
        case 'shutdown-hard':
            changeStatus('connected-inactive', "hard shutdown triggered")
            break
    }
}

/**
 * modification of https://stackoverflow.com/a/79970
 * @param {HTMLButtonElement} btn 
 * @param {string} cmdname - see cmd parameter of sendMasterCommand
 * @param {number} dur
 */
function execHold(btn, cmdname, dur) {
    var t;
    btn.onmousedown = function () {
        setTimeout(function () {
            sendMasterCommand(cmdname)
        }, dur);
    }
    btn.onmouseup = function () {
        clearTimeout(t);
    }
};