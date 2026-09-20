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
	_ "embed"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed LICENSE
var gplLic string

func main() {
	log.Writer()
	masconfPath := flag.String("c", "", "location of master config")
	noweb := flag.Bool("noweb", false, "don't load web interface")
	ginDbg := flag.Bool("gin-debug", false, "set gin debug")
	dummyEstop := flag.Bool("dummy-estop", false, "dummy e-stop")

	help := flag.Bool("help", false, "show long help (recommended)")
	lic := flag.Bool("license", false, "show the full GPLv3 license this program is under")
	flag.Parse()

	if *help {
		fmt.Fprintln(os.Stderr, `broadcastinator: youtube-oriented RTMP broadcast thingy

Options:
	-c <configfile>    path to the configuration file
	--noweb            don't load web interface
	--render           render out the entire channel into a single video file
	--noprelive-test   don't check if channel and its associated programs are ready to use
	--gin-debug        enable Gin's debug log
	--dummy-estop      make /masterctl/e-stop do everything exept exit the program
	--help             show this long help`)
		return
	} else if *lic {
		fmt.Fprintln(os.Stderr, gplLic)
		return
	}

	if !*ginDbg {
		gin.SetMode(gin.ReleaseMode)
	}

	masconf, err := os.Open(*masconfPath)
	chk(err)
	config := MasterConfig{}
	config = parseGeneric(masconf, config)

	var logf *os.File
	if config.LogFilePath != "" {
		logf = openOrCreate(config.LogFilePath)
		globalLogger = slog.New(slog.NewTextHandler(logf, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		globalLogger.Info(fmt.Sprintf("----- PROGRAM STARTED (%s) -----", time.Now().Format(time.RFC1123Z)))
	}

	if !*noweb {
		fmt.Fprintln(os.Stderr, "Loading web interface")
		globalLogger.Info("Loading web interface")
		go loadInterface()
	}
	fmt.Fprintln(os.Stderr, "Loading API")
	globalLogger.Info("Loading APIs")
	go loadAPI(logf, *dummyEstop)

	// now we wait...
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	globalLogger.Info(fmt.Sprintf("----- PROGRAM STOPPED (%s) -----", time.Now().Format(time.RFC1123Z)))
}
