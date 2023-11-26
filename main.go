package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"telegram-bot/menu"
	"telegram-bot/tconfig"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/telebot.v3"
)

var db *sql.DB
var rdb *redis.Client

func main() {
	config, err := tconfig.ParseConfig("etc/config.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = initDb(config)
	if err != nil {
		log.Fatal(err)
	}

	err = initCache(config)
	if err != nil {
		log.Fatal(err)
	}

	pref := telebot.Settings{
		Token:  config.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	tele, err := telebot.NewBot(pref)

	// tele.Use(middleware.Logger())
	if err != nil {
		log.Fatal(err)
		return
	}

	menu.LoadAllMenu(tele)

	tele.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send(strconv.FormatInt(c.Chat().ID, 10) + c.Chat().Username)
	})

	// tags
	tele.Handle("/tags", func(c telebot.Context) error {
		tags := c.Args()
		args := ""
		for _, tag := range tags {
			args += tag
		}
		return c.Send(args)
	})

	// payload
	tele.Handle("/clean", func(c telebot.Context) error {
		fmt.Println()
		return nil
	})

	tele.Start()
}

func initDb(config tconfig.Tconfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Mysql.Username, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.Database)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil
}

func initCache(config tconfig.Tconfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
