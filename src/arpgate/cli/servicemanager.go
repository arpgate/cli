package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const SERVICE_STRONGSWAN = "ipsec"
const SERVICE_DHCCP = "isc-dhscp-server"
const SERVICE_DNS = "bind9"
const SERVICE_MQTT = "mqtt"

const SERVICE_TFTP = "tftpd-hpa"

const SERVICE_HAPROXY = "haproxy"
const SERVICE_NFINX = "nginx"
const SERVICE_SNORT = "snort"

func RestartService(name string) {
	cmd := exec.Command("service", name, "restart")	
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
