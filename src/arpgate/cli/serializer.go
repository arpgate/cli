package cli

import (
	"gopkg.in/yaml.v2"
	"log"
	"time"
)

func BuildYaml(conf Configuration) string {
	conf.Created = string(time.Now().Format(time.RFC3339))
	conf.Release = RELEASE
	conf.SoftwareVersion = VERSION
	return Serialize(conf)
}

func Serialize(conf Configuration) string {
	y, err := yaml.Marshal(conf)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(y)
}

func DeSerialize(conf string) Configuration {
	out := Configuration{}
	yaml.Unmarshal([]byte(conf), &out)
	return out
}
