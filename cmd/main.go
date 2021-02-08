package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`./config.json`)
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Connont find config file, %s", err)
	}
}

func main() {
	// code goes here
}

func checkError(err error) {
	if nil != err {
		log.Printf("error: %v\n", err)
		panic(err)
	}
}
