package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type ROOM map[string]PC

type PC struct {
	Name      string `yaml:"name"`
	WorkGroup string `yaml:"group"`
	Netmask   string `yaml:"nm"`
	Gateway   string `yaml:"gw"`
	DNS       string `yaml:"dns"`
	IP        string `yaml:"ip"`
}

var (
	room    ROOM                // mac = {}
	rooms   = map[string]ROOM{} // room = {}
	roomName = ""
)

func parseRoom(rom string) (ROOM, error) {
	roomPath := filepath.Join("rooms", rom+".yml")
	f, err := os.ReadFile(roomPath)
	if err != nil {
		return ROOM{}, err
	}
	err = yaml.Unmarshal(f, &room)
	if err != nil {
		return ROOM{}, err
	}
	return room, nil
}

func loadRooms() {
	filepath.WalkDir("rooms", func(p string, d fs.DirEntry, er error) error {
		if strings.Contains(p, ".yml") {
			roomName = strings.TrimSuffix(filepath.Base(p), filepath.Ext(p))
			rom, err := parseRoom(roomName)
			delete(rom, "define")
			rooms[roomName] = rom
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		return nil
	})
}
