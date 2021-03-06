// XML
// go.format.xml.test.go
//

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIp string `xml:"serverIP"`
}

type Servers struct {
	Name xml.Name `xml:"servers"`
	Version int `xml:"version"`
	Servers []Server `xml:"server"`
}

func main() {
	data, err := ioutil.ReadFile("./my.xml")
	if err != nil {
		fmt.Println("ioutil.ReadFile err : ", err)
		return
	}

	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("xml: %v\n", servers)
}
