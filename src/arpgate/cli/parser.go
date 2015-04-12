package cli

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const IPESCCONF = "/etc/ipsec.conf"
const IPSECSECRETS = "/etc/ipsec.secrets"

const STRONGSWAN_SHAREDSECRET = "{STRONGSWAN_SHAREDSECRET}"
const STRONGSWAN_USER = "{STRONGSWAN_USER}"
const STRONGSWAN_PWD = "{STRONGSWAN_PWD}"

func ParseStrongSwan(conf Configuration) {
	ipscsecrets := loadTemplate("ipsec.secrets")

	fmt.Println(ipscsecrets)
	fmt.Println(conf.Vpn.Sharedsecret)

	ipscsecrets = strings.Replace(ipscsecrets, STRONGSWAN_SHAREDSECRET, conf.Vpn.Sharedsecret, 1)
	ipscsecrets = strings.Replace(ipscsecrets, STRONGSWAN_USER, conf.Vpn.Username, 1)
	ipscsecrets = strings.Replace(ipscsecrets, STRONGSWAN_PWD, conf.Vpn.Pwd, 1)

	fmt.Println(ipscsecrets)

	d1 := []byte(ipscsecrets)
	err := ioutil.WriteFile(IPSECSECRETS, d1, 0644)
	check(err)
}

func loadTemplate(name string) string {
	dat, err := ioutil.ReadFile(TEMPLATEFOLDER + name)
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
