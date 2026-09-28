# Web interface

As of this writing the web interface is very simple.

On the bottom you see 5 buttons. All except one labeled "Emergency Stop" need to be held down for 2 seconds ("Emergency Stop" needs to be held for 5 seconds)

**Start**, normally, starts the livestream. If however the `start` property on the master configuration is set, the livestream will start with a screen specified in the assets configuration file. Once the time has passed, then will the channel actually start broadcasting.

**Override** forces any time-delayed action to run. This is used for example to force a channel to start broadcasting regardless if `start` specifies a time or not.

**Shutdown** has two buttons, one for a "soft" shutdown and one for a "hard" shutdown. "soft" means there will be a "goodbye" card shown, usually reserved when the last program ends. "hard" means the broadcast will suddenly stop


**Emergency Stop** essentially exits the entire program (the actual program not the livestream) from the web console. *Once this has triggered, reloading the web interface will result in a blank page.*