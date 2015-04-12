package clitest

import (
	"arpgate/cli"
	"fmt"
)

func Runtest() {

	// testing serializer
	yamlStr := cli.BuildYaml(cli.GetDefault())
	fmt.Println(yamlStr)

	config := cli.DeSerialize(yamlStr)
	fmt.Println(config)
}
