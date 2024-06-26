package httpRequest

import (
	"fmt"
	"net"
	"strings"
)

func HttpTCPConnection(host, port string) (net.Conn, error) {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PrepareRequest(requestFlag string, headers []string, dataFlag string, host string, path string) (string, error) {
	method := strings.ToUpper(requestFlag)
	request := fmt.Sprintf("%s %s HTTP/1.1\n", method, path)
	request += fmt.Sprintf("Host: %s\n", host)
	request += "Accept: */*\n"
	request += "Connection: close\n"

	for _, header := range headers {
		request += header + "\n"
	}

	if dataFlag != "" {
		request += fmt.Sprintf("Content-Length: %d\n", len(dataFlag))
		request += "\n"
		request += dataFlag
	}

	request += "\n"

	return request, nil
}
