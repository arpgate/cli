package cli

func GetDefault() Configuration {

	conf := Configuration{}

	// device
	device := Device{}
	device.Email = ""
	device.Ip = "10.0.0.253"
	device.Gateway = "10.0.0.1"
	device.Netmask = "255.255.255.0"

	conf.Device = device

	// Dhcp
	conf.Dhcp.Enabled = false
	conf.Dhcp.Pxebootenabled = true
	conf.Dhcp.Rangelow = "10.0.0.100"
	conf.Dhcp.Rangehigh = "10.0.0.200"

	// dns
	rec := DnsRecord{"A", "pc", "10.0.0.2"}
	conf.Dns.DnsRecords = make([]DnsRecord, 1)
	conf.Dns.DnsRecords[0] = rec
	conf.Dns.ForwarderA = "8.8.8.8"
	conf.Dns.ForwarderB = "8.8.4.4"
	conf.Dns.Enabled = true

	// vpn
	conf.Vpn.Enabled = false
	conf.Vpn.Username = "arpgate"
	conf.Vpn.Pwd = "changme"
	conf.Vpn.Sharedsecret = "changme_to_a_secret_secret"

	// ntp
	conf.Ntp.Enabled = true
	conf.Ntp.ServerA = "0.north-america.pool.ntp.org"
	conf.Ntp.ServerB = "1.north-america.pool.ntp.org"

	// haproxy
	conf.Haproxy.Enabled = true
	nodes := []Node{Node{"127.0.0.1", 8887}}
	forwarding := Forwarding{8888, nodes}
	conf.Haproxy.Forwardings = make([]Forwarding, 1)
	conf.Haproxy.Forwardings[0] = forwarding

	// nginx
	conf.Nginx.Enabled = true
	site := Site{"*", "/var/www/arpgate", 8887, true}
	conf.Nginx.Sites = make([]Site, 1)
	conf.Nginx.Sites[0] = site

	// mqtt
	conf.Mqtt.Enabled = true

	// snort
	conf.Snort.Enabled = false
	return conf

}
