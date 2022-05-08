package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	tr "github.com/snakesel/libretranslate"
)

func main() {
	engFile := flag.String("base", "_locales/en/messages.json", "Messages to generate translations for. English to all other supported lanugages only for now.")
	flag.Parse()
	jsonBytes, err := ioutil.ReadFile(*engFile)
	if err != nil {
		panic(err)
	}
	var x map[string]interface{}

	translate := tr.New(tr.Config{
		Url: "https://libretranslate.com",
		Key: "XXX",
	})

	json.Unmarshal(jsonBytes, &x)
	for _, v := range x {
		kv := v.(map[string]interface{})
		for ik, iv := range kv {
			fmt.Println(ik, ":", iv)
			// you can use "auto" for source language
			// so, translator will detect language
			trtext, err := translate.Translate(iv, "auto", "ru")
			if err == nil {
				fmt.Println(trtext)
			} else {
				fmt.Println(err.Error())
			}
		}
	}
}
