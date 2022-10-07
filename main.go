package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/eyedeekay/goSam"
	tr "github.com/snakesel/libretranslate"
)

/*
English en
Arabic ar
Chinese zh
French fr
German de
Italian it
Japanese jp
Portuguese pr
Russian ru
Spanish es
*/

var langs = []string{
	"en",
	"ar",
	"zh",
	"fr",
	"de",
	"it",
	"ja",
	"pt",
	"ru",
	"es",
}

func main() {
	engFile := flag.String("base", "_locales/en/messages.json", "Messages to generate translations for")
	lang := flag.String("lang", "", "Select language to translate to(default will translate all)")
	flag.Parse()
	if *lang == "" {
		locks := len(langs)
		for _, lang := range langs {
			transDir := filepath.Join("_locales", lang)
			transFile := filepath.Join(transDir, "messages.json")
			if transFile != *engFile {
				go func() {
					os.MkdirAll(transDir, 0755)
					jsonBytes, err := ioutil.ReadFile(*engFile)
					if err != nil {
						panic(err)
					}
					var x map[string]interface{}

					sam, err := goSam.NewDefaultClient()
					if err != nil {
						panic(err)
					}
					defer sam.Close()
					//checkErr(err)

					log.Println("Client Created")

					// create a transport that uses SAM to dial TCP Connections
					httpClient := &http.Client{
						Transport: &http.Transport{
							Dial: sam.Dial,
						},
					}

					http.DefaultClient = httpClient

					translate := tr.New(tr.Config{
						Url: "http://w62j277kjls7agmctbtjzuthvsaiz7zzjthmahdk7pgweditlfzq.b32.i2p",
					})
					json.Unmarshal(jsonBytes, &x)

					y := make(map[string]interface{})
					for k, v := range x {
						y[k] = v
						kv := v.(map[string]interface{})
						for ik, iv := range kv {
							//fmt.Println(ik, ":", iv)
							// you can use "auto" for source language
							// so, translator will detect language
							trtext, err := translate.Translate(iv.(string), "auto", lang)
							if err != nil {
								fmt.Println(err.Error())
								if strings.Contains(err.Error(), "Slowdown") {
									time.Sleep(time.Minute + time.Second)
									trtext, err = translate.Translate(iv.(string), "auto", lang)
									if err != nil {
										panic(err)
									}
								} else {
									panic(err)
								}
							}
							z := y[k].(map[string]interface{})
							z[ik] = trtext
							y[k] = z
							//fmt.Println(trtext)
							fmt.Println(z)
						}
						jsonStr, err := json.MarshalIndent(y, "", "    ")
						if err != nil {
							panic(err)
						}
						err = ioutil.WriteFile(transFile, []byte(jsonStr), 0644)
						if err != nil {
							panic(err)
						}
					}
					locks--
				}()
			}
			for locks > 0 {
				fmt.Println("waiting for:", locks, "jobs to complete")
			}
		}
	} else {
		transDir := filepath.Join("_locales", *lang)
		transFile := filepath.Join(transDir, "messages.json")
		if transFile != *engFile {
			os.MkdirAll(transDir, 0755)
			jsonBytes, err := ioutil.ReadFile(*engFile)
			if err != nil {
				panic(err)
			}
			var x map[string]interface{}

			sam, err := goSam.NewDefaultClient()
			if err != nil {
				panic(err)
			}
			defer sam.Close()
			//checkErr(err)

			log.Println("Client Created")

			// create a transport that uses SAM to dial TCP Connections
			httpClient := &http.Client{
				Transport: &http.Transport{
					Dial: sam.Dial,
				},
			}

			http.DefaultClient = httpClient

			translate := tr.New(tr.Config{
				Url: "http://w62j277kjls7agmctbtjzuthvsaiz7zzjthmahdk7pgweditlfzq.b32.i2p",
			})
			json.Unmarshal(jsonBytes, &x)

			y := make(map[string]interface{})
			for k, v := range x {
				y[k] = v
				kv := v.(map[string]interface{})
				for ik, iv := range kv {
					//fmt.Println(ik, ":", iv)
					// you can use "auto" for source language
					// so, translator will detect language
					trtext, err := translate.Translate(iv.(string), "auto", *lang)
					if err != nil {
						fmt.Println(err.Error())
						if strings.Contains(err.Error(), "Slowdown") {
							time.Sleep(time.Minute + time.Second)
							trtext, err = translate.Translate(iv.(string), "auto", *lang)
							if err != nil {
								panic(err)
							}
						} else {
							panic(err)
						}
					}
					z := y[k].(map[string]interface{})
					z[ik] = trtext
					y[k] = z
					//fmt.Println(trtext)
					fmt.Println(z)
				}
				jsonStr, err := json.MarshalIndent(y, "", "    ")
				if err != nil {
					panic(err)
				}
				err = ioutil.WriteFile(transFile, []byte(jsonStr), 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
