package app

import (
	"log"
	"os"
)

func CreateDirStaticIsNotExist() {
	if _, err := os.Stat("static"); err != nil {
		err = os.Mkdir("static", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		err = os.Mkdir("static/images", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		err = os.Mkdir("static/pdfs", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
