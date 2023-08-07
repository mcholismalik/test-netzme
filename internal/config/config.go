package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	App        App        `json:"app"`
	Swagger    Swagger    `json:"swagger"`
	Repository Repository `json:"repository"`
}

var Config *Configuration = &Configuration{}

func Load(path string) (*Configuration, error) {
	if path == "" {
		wd, err := os.Getwd()
		if err != nil {
			return Config, err
		}

		env := os.Getenv("ENV")
		if env == "" {
			env = "local"
		}
		path = fmt.Sprintf("%s/config/config.%s.yml", wd, env)
	}

	err := configor.New(&configor.Config{AutoReload: true, AutoReloadInterval: time.Minute}).Load(Config, path)
	if err != nil {
		logrus.Info(err)
		return Config, err
	}

	return Config, nil
}

func (Configuration) String() string {
	sb := strings.Builder{}
	return sb.String()
}

func (c *Configuration) Raw() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}
