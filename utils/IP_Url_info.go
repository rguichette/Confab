package utils

import (
	"net"
)

type IpDetailGrabber struct {
	Grabber IpInfoGrabber
}

type IpInfoGrabber interface {
	GetServerInfo() *IpDetails
}

type IpDetails struct {
	ServerIp   string
	ServerPort string
	ServerUrl  string
}

func (ip *IpDetails) GetServerInfo() *IpDetails {
	serverIp := getLocalIP()
	serverPort := "9999"
	serverUrl := "http://" + serverIp + ":" + serverPort

	return &IpDetails{
		serverIp,
		serverPort,
		serverUrl,
	}

}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "unknown"
}

func NewIpDetailsGrabber() IpDetailGrabber {

	return IpDetailGrabber{
		Grabber: &IpDetails{},
	}

}
