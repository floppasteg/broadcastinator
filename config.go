package main

import (
	"io"
	"time"

	"github.com/BurntSushi/toml"
)

type MasterConfig struct {
	StreamURL        string `toml:"streamurl"`
	StreamKeyFile    string `toml:"streamkey"`
	ServiceSpecifier string `toml:"service-spec"`
	ProgramRoot      string `toml:"program"`
	ExtraAssets      string `toml:"extras"`
	LogFilePath      string `toml:"logfile"`
}

// Used in programs

type ChannelProgram struct {
	Title      string        `toml:"title"`
	ConfigPath string        `toml:"conf"`
	Runtime    time.Duration `toml:"runtime"`
}

type ChannelConfig struct {
	Programs          []ChannelProgram `toml:"programs"`
	ShowScheduleAfter int              `toml:"show-schedule-after"`
}

type Break struct {
	VideoTimestamp    string `toml:"timestamp"`
	ShowTitleCard     bool   `toml:"show-title-card"`
	ShowYoureWatching bool   `toml:"show-youre-watching"`
}

type ProgramConfig struct {
	VideoFile          string   `toml:"video"`
	Ratings            []string `toml:"ratings"`
	Breaks             []Break  `toml:"breaks"`
	BreakTitleCard     string   `toml:"break-title-card"`
	BreakYoureWatching string   `toml:"break-youre-watching"`

	confpath string
}

// Used in asset files

type AssetConfig struct {
	InterruptCard string `toml:"interrupt"`
	EASCard       string `toml:"eas"`
}

func parseGeneric[T MasterConfig | ChannelConfig](confile io.Reader, conf T) T {
	dec := toml.NewDecoder(confile)
	_, err := dec.Decode(&conf)
	chk(err)
	return conf
}
