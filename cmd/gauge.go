package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lukemilby/gauge/pkg/gauge"
	"github.com/spf13/viper"
)

var i = flag.String("i", "5D", "interval time for market range")

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	flag.Parse()

	key := viper.GetString("APIKEY")
	// create new gauge instance
	g := gauge.NewGauge(gauge.Config{
		Interval: *i,
		ApiKey:   key,
	})

	// TODO: Print to stdout
	fmt.Println(g.Run())
}
