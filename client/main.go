package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type HostInfo struct {
	NamePrefix string
	IPPrefix   string
	NetMask    string
	GateWay    string
	DNS1       string
	DNS2       string
}

var (
	hi           HostInfo
	card         net.Interface
	computerName string
	ip           string
	num          string
	inum         uint64
	err          error
	argsize      int
	room         string
)

func parseArgs() {
	argsize = len(os.Args)
	if argsize != 3 {
		log.Fatalln("参数不对, 需要2个参数,第1为电脑室编号,第2为电脑编号")
	}

	num = os.Args[2]
	inum, err = strconv.ParseUint(num, 10, 32)

	if err != nil {
		log.Fatalln("必须是数字")
	}
	if inum < 1 {
		log.Fatalln("数字必须大于1")
	}

	room = os.Args[1]
	hi.DNS1 = "10.2.104.18"
	hi.DNS2 = "114.114.114.114"
	hi.GateWay = fmt.Sprintf("192.168.%s.254", room)
	hi.IPPrefix = fmt.Sprintf("192.168.%s.", room)
	hi.NetMask = "255.255.255.0"
	hi.NamePrefix = "st"
}

func makeHost() {
	ip = hi.IPPrefix + num
	if inum < 10 {
		computerName = fmt.Sprintf("%s0%d", hi.NamePrefix, inum)
	} else {
		computerName = fmt.Sprintf("%s%d", hi.NamePrefix, inum)
	}
}

// func parseToml() {
// 	_, err := toml.DecodeFile("hostinfo.toml", &hi)
// 	if err != nil {
// 		log.Fatalln("不能有解析配置文件")
// 	}
// }

func init() {
	parseArgs()
	// parseToml()
}

func main() {
	findCard()
	makeHost()
	setNic()
	setName()
}

func findCard() {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if inter.Name == "本地连接" {
			card = inter
			break
		} else {
			continue
		}
	}
}

func setNic() {
	if room == "0" {
		dhcp := fmt.Sprintf("netsh interface ip set address %s dhcp", card.Name)
		mycmd(dhcp)
		dns := fmt.Sprintf("netsh interface ip set dns name=%s dhcp", card.Name)
		mycmd(dns)
		return
	}
	cmdIP := fmt.Sprintf("netsh interface ip set address name=%s static %s %s %s", card.Name, ip, hi.NetMask, hi.GateWay)
	cmdDNS1 := fmt.Sprintf("netsh interface ip set dns name=%s static %s", card.Name, hi.DNS1)
	cmdDNS2 := fmt.Sprintf("netsh interface ip add dns name=%s %s index=2", card.Name, hi.DNS2)
	mycmd(cmdIP)
	mycmd(cmdDNS1)
	mycmd(cmdDNS2)
}

func setName() {
	oldName, _ := exec.Command("cmd", "/C", "hostname").Output()
	nameStr := strings.TrimSuffix(string(oldName), "\r\n")
	cmdName := fmt.Sprintf("WMIC computersystem where caption='%s' rename %s", nameStr, computerName)
	err = mycmd(cmdName)
	if err != nil {
		log.Println(err.Error())
	}
}

func mycmd(str string) error {
	c := exec.Command("cmd","/C",str)
	return c.Run()
}