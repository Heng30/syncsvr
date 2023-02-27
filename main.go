package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"local/db"
	"local/svr"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	ListenAddr string
	TestMode   bool
}

const DB_BAME string = "syncsvr.db"

var confPath, dbPath string
var appConf Config = Config{ListenAddr: ":8000", TestMode: false}

var (
    queryToken = flag.Bool("query", false, "query all access tokens")
	addToken = flag.String("add", "", "add access token")
	delToken = flag.String("del", "", "delete access token")
	updToken = flag.String("upd", "", "update access token")
)

func main() {
	parseArgs()
	initEnv()
	db.Init(dbPath + "/" + DB_BAME)
	defer db.Uninit()

	if flag.NFlag() > 0 {
		runAsCln()
	} else {
		runAsSvr()
	}
}

func parseArgs() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-add <token>] | [del <token>] | [upd <oldToken,newToken>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Version:", VERSION)
		fmt.Println("Configure: ~/.config/syncsvr/config.json")
		fmt.Println("Database: ~/.local/share/syncsvr/syncsvr.db")
	}
	flag.Parse()
}

func runAsCln() {
    if *queryToken {
        if tokens, err := db.QueryAccessTokens(); err != nil {
            log.Fatalln(err)
        } else {
            fmt.Println(tokens)
        }
    } else if *addToken != "" {
		if err := db.AddAccessToken(*addToken); err != nil {
			log.Fatalln(err)
		}
	} else if *delToken != "" {
		if err := db.DelAccessToken(*delToken); err != nil {
			log.Fatalln(err)
		}
	} else if *updToken != "" {
		tokens := strings.Split(*updToken, ",")
		if len(tokens) != 2 {
			log.Fatalln("Invalid format:", *updToken)
		}
        if err := db.UpdateTableName(tokens[0], tokens[1]); err != nil {
			log.Fatalln(err)
        }
		if err := db.UpdateAccessToken(tokens[0], tokens[1]); err != nil {
			log.Fatalln(err)
		}
	}
}

func runAsSvr() {
	svr.Start(appConf.ListenAddr, appConf.TestMode)
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
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	if err := enc.Encode(appConf); err != nil {
		log.Fatalln("Error:", err)
	}
}
