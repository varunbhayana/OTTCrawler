package conf

import (
	"encoding/json"
	"io/ioutil"
)

// parse from conf.json
type confjson struct {
	Queues map[string]interface{}
}

const (
	confFile = "./conf/conf.json"
)

var (
	D confjson
)

func init() {
	cfile, err := ioutil.ReadFile(confFile)

	if err != nil {
		panic("fail to load configuration file")
	}

	if err := json.Unmarshal(cfile, &D); err != nil {
		panic(err)
	}

}
