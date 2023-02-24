package main

import (
	"encoding/json"
	"local/db"
	"local/svr"
	"log"
	"os"
)

type Config struct {
	Tokens []string
}

var confPath, dbPath string
var appConf Config = Config{Tokens: []string{}}

func main() {
	log.Println("start...")
	initEnv()
	db.Init(dbPath)
	svr.Start(":8080")
	log.Fatalln("exit...")
}

func initEnv() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	home, e := os.LookupEnv("HOME")
	if !e {
		log.Fatalln("can't find $HOME")
	}

	confPath = home + "/.config/syncsvr"
	dbPath = home + "/.local/share/syncsvr"
	for _, path := range []string{confPath, dbPath} {
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(path, os.ModePerm); err != nil {
					log.Fatalln("Error:", err)
				}
			} else {
				log.Fatalln("Error:", err)
			}
		}
	}
	loadConf()
}

func loadConf() {
	config := confPath + "/config.json"
	if _, err := os.Stat(config); err != nil {
		if os.IsNotExist(err) {
			defaultConf(config)
		} else {
			log.Fatalln("Error:", err)
		}
	} else {
		if js, err := os.ReadFile(config); err != nil {
			log.Fatalln("Error:", err)
		} else {
			if err := json.Unmarshal(js, &appConf); err != nil {
				log.Fatalln("Error:", err)
			}
		}
	}
}

func defaultConf(path string) {
    file, err := os.OpenFile(path, os.O_CREATE | os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalln("Error:", err)
    }
    defer file.Close()

    enc := json.NewEncoder(file)
    if err := enc.Encode(appConf); err != nil {
        log.Fatalln("Error:", err)
    }
}
