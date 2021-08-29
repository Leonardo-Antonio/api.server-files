package helper

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func SaveFile(fileHeader *multipart.FileHeader, path string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	fileName := id + "__" + strings.Join(strings.Split(fileHeader.Filename, " "), "-")

	err = ioutil.WriteFile(
		path+fileName,
		fileBytes,
		os.ModePerm,
	)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
