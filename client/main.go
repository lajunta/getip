package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
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
	inum         int
	err          error
)

func parseArgs() {
	size := len(os.Args)
	if size != 2 {
		log.Fatalln("参数不对, 使用: getip 3")
	}
	num = os.Args[1]
	inum, err = strconv.Atoi(num)
	if err != nil {
		log.Fatalln("必须是数字")
	}
	if inum < 1 {
		log.Fatalln("数字必须大于1")
	}
}

func makeHost() {
	ip = hi.IPPrefix + num
	if inum < 10 {
		computerName = fmt.Sprintf("%s0%d", hi.NamePrefix, inum)
	} else {
		computerName = fmt.Sprintf("%s%d", hi.NamePrefix, inum)
	}
}

func parseToml() {
	_, err := toml.DecodeFile("hostinfo.toml", &hi)
	if err != nil {
		log.Fatalln("不能有解析配置文件")
	}
}

func init() {
	parseArgs()
	parseToml()
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
	cmdIP := fmt.Sprintf("netsh interface ip set address name=%s static %s %s %s", card.Name, ip, hi.NetMask, hi.GateWay)
	cmdDNS1 := fmt.Sprintf("netsh interface ip set dns name=%s static %s", card.Name, hi.DNS1)
	cmdDNS2 := fmt.Sprintf("netsh interface ip add dns name=%s %s index=2", card.Name, hi.DNS2)
	cmd := exec.Command("cmd", "/C", cmdIP)
	cmd.Run()
	cmd = exec.Command("cmd", "/C", cmdDNS1)
	cmd.Run()
	cmd = exec.Command("cmd", "/C", cmdDNS2)
	cmd.Run()
}

func setName() {
	oldName, _ := exec.Command("cmd", "/C", "hostname").Output()
	nameStr := strings.TrimSuffix(string(oldName), "\r\n")
	cmdName := fmt.Sprintf("WMIC computersystem where caption='%s' rename %s", nameStr, computerName)
	cmd := exec.Command("cmd", "/C", cmdName)
	err = cmd.Run()
	if err != nil {
		log.Println(err.Error())
	}

}
