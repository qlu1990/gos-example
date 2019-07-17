package conf

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//Mongodb struct  mongodb conf
type Mongodb struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

//Server struct server mongodb conf
type Server struct {
	Address string `yaml:"address"`
}

//Conf struct for conf file
type Conf struct {
	Mongodb Mongodb `yaml:"mongodb"`
	Server  Server  `yaml:"server"`
}

//Cfg global var conf
var Cfg Conf

//LoadConf load conf file
func LoadConf() {
	workdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := ioutil.ReadFile(filepath.Join(workdir, "conf.yml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, &Cfg)
	if err != nil {
		panic(err)
	}
}
