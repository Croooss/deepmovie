package config

import (
	"os"
	"path/filepath"
	log "github.com/cihub/seelog"
	"fmt"
	"os/exec"
	"io/ioutil"
	"encoding/json"
)

type DataFile struct {
	Movie string
	Rating string
	Tag string
}


type AppConfig struct {
	Files *DataFile
	UserID string
	Count int
}

func init() {
	logInit()
}

func logInit() {
	logger, err := log.LoggerFromConfigAsFile(LoadConfigDir() + DefaultLogConfigFile)
	if err != nil {
		fmt.Printf("log config file load err : %v \n" , err)
		os.Exit(1)
	}

	log.ReplaceLogger(logger)
}


func LoadCurDir() string {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Errorf("application path error:%v \n", err)
		os.Exit(1)
	}

	appPath, _ := filepath.Abs(path)
	curdir, _ := filepath.Split(appPath)

	return curdir
}

func LoadConfigDir() string {
	dir :=  LoadCurDir() + "bin/conf"
	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModePerm)
	}

	return dir
}

func LoadConfig(file string, conf interface{}) error {
	fileData, err := ioutil.ReadFile(LoadConfigDir() + "/" + file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, conf)
	if err != nil {
		return err
	}

	return nil
}