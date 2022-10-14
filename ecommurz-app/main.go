package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type conf struct {
	Host     string `yaml:"DB_HOST"`
	Port     int32  `yaml:"DB_PORT"`
	User     string `yaml:"DB_USERNAME"`
	Password string `yaml:"DB_PASSWORD"`
	Name     string `yaml:"DB_NAME"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("config/conf.yaml")
	if err != nil {
		log.Error().Stack().Err(errors.New(err.Error())).Msg("")
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Error().Stack().Err(errors.New(err.Error())).Msg("")
	}
	return c
}

func connectToDB() (bool, conf) {
	var c conf
	c.getConf()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Error().Stack().Err(errors.New(err.Error())).Msg("")
		return false, c
	}
	return true, c
}

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, request *http.Request) {
			isConnected, c := connectToDB()
			// Try connecting to db every 2s
			for !isConnected {
				time.Sleep(2 * time.Second)
				isConnected, c = connectToDB()
			}
			log.Debug().Msg(fmt.Sprintf("successfully connected"))
			log.Debug().Msg(fmt.Sprintf("current username: %s", c.User))
			log.Debug().Msg(fmt.Sprintf("current password: %s", c.Password))
		})
	http.ListenAndServe(":8090", nil)
}
