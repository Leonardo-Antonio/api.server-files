package helper

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func SaveFile(fileHeader *multipart.FileHeader, path, typeFile string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	id, err := gonanoid.New(8)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(
		path+fmt.Sprintf("%s.%s", id, typeFile),
		fileBytes,
		os.ModePerm,
	)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", id, typeFile), nil
}
