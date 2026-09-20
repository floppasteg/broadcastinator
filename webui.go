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
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func loadInterface() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("ui"))
	mux.Handle("/", fs)
	chk(http.ListenAndServe(":8080", mux))
}

const masterctlAPIRoot = "/masterctl"

func loadAPI(logfile *os.File, dummyexit bool) {
	route := gin.Default()
	route.Use(gin.LoggerWithWriter(logfile))

	// because of browser security we need to add CORS headers even if everything here is localhost
	route.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"*"},
		AllowOrigins: []string{"http://localhost:8080", "http://127.0.0.1:8080"},
	}))

	route.POST(masterctlAPIRoot+"/start", func(ctx *gin.Context) {
		globalLogger.Debug("console requested to start livestreaming")
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "",
			"timestamp": time.Now().UnixMilli(),
		})
	})
	route.POST(masterctlAPIRoot+"/shutdown-soft", func(ctx *gin.Context) {
		globalLogger.Debug("console requested a soft shutdown")
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "",
			"timestamp": time.Now().UnixMilli(),
		})
	})
	route.POST(masterctlAPIRoot+"/shutdown-hard", func(ctx *gin.Context) {
		globalLogger.Debug("console requested a hard shutdown")
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "",
			"timestamp": time.Now().UnixMilli(),
		})
	})
	route.POST(masterctlAPIRoot+"/e-stop", func(ctx *gin.Context) {
		globalLogger.Debug("console requested a *emergency stop*")
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "",
			"timestamp": time.Now().UnixMilli(),
		})
		if !dummyexit {
			os.Exit(0)
		}
	})
	route.GET("/conn", func(ctx *gin.Context) {
		globalLogger.Debug("connection test requested")
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "no data yet",
			"timestamp": time.Now().UnixMilli(),
		})
	})
	route.Run(":8088")
}
