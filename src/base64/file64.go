package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
    "flag"
	"io"
	"os"
	"strings"
)

func Encode(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buff), nil
}

func Decode(code string, dest string) error {
	buff, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dest, buff, 07440); err != nil {
		return err
	}

	return nil
}

func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}

func main() {
    flag.Parse()
    args := flag.Args()

    pdf := args[0]
    log.Print(pdf)
    
	dataStr, _ := Encode(pdf)
	WriteStringToFile(pdf + ".txt", dataStr)
}
