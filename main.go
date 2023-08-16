package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var (
	KeyChannel chan string
)

func init() {
	logo := "ooooo   ooooo                                             ooooooooo.   ooooooooooooo \n" +
		"`888'   `888'                                             `888   `Y88. 8'   888   `8 \n" +
		" 888     888   .oooo.   oo.ooooo.  oo.ooooo.  oooo    ooo  888   .d88'      888      \n" +
		" 888ooooo888  `P  )88b   888' `88b  888' `88b  `88.  .8'   888ooo88P'       888      \n" +
		" 888     888   .oP\"888   888   888  888   888   `88..8'    888`88b.         888      \n" +
		" 888     888  d8(  888   888   888  888   888    `888'     888  `88b.       888      \n" +
		"o888o   o888o `Y888\"\"8o  888bod8P'  888bod8P'     .8'     o888o  o888o     o888o     \n" +
		"                         888        888       .o..P'                                 \n" +
		"                        o888o      o888o      `Y8P'  \n" +
		"								by  Bored Monkey 	 \n"
	print(logo)

}
func main() {
	FilePath := flag.String("f", "", "字典目录")
	Times := flag.Int("t", 100, "快乐次数")
	DnslogDomain := flag.String("d", "", "dnslog域名")
	flag.Parse()
	if *DnslogDomain == "" {
		fmt.Println("请输入快乐配送地址")
		return
	}
	if *FilePath == "" {
		KeyChannel = make(chan string, 20)
		KeyChannel <- "root"
		KeyChannel <- "administrator"
		KeyChannel <- "WIN-2K3SERVERX64"
		KeyChannel <- "administrator/WIN-2K3SERVERX86"
		KeyChannel <- "administrator/DC-01"
		KeyChannel <- "nginx-01"
		KeyChannel <- "nginx-02"
		KeyChannel <- "Control-HuBei"
		KeyChannel <- "DC-Peiking"
		KeyChannel <- "31032"
		KeyChannel <- "55"
		KeyChannel <- "ens33"
		KeyChannel <- "ens32"
		KeyChannel <- time.Now().String()
		KeyChannel <- "TSD-01"
		KeyChannel <- "TSD-02"
		KeyChannel <- "TLMSHMESBTDBB"
		KeyChannel <- "EPicp.cpnc.com.cn"
		KeyChannel <- "ERP"
		KeyChannel <- "CRM"
	} else {
		Dir, err := os.ReadFile(*FilePath)
		if err != nil {
			fmt.Println("读取字典失败")
			return
		}
		DirArray := strings.Split(string(Dir), "\n")
		KeyChannel = make(chan string, len(DirArray))
		for _, v := range DirArray {
			KeyChannel <- v
		}

	}
	if *Times == -1 {
		RateControllor := make(chan int, 2000)
		fmt.Println("开始究极快乐")
		i := 0
		for {
			RateControllor <- 1
			go func() {
				Sub := <-KeyChannel
				net.LookupIP(Sub + "." + *DnslogDomain)
				KeyChannel <- Sub
				i++
				if i%10 == 0 {
					fmt.Println("红队究极快乐中...，快乐次数：", i)
				}
				<-RateControllor
			}()

		}
	} else {
		for i := 0; i < *Times; i++ {
			Sub := <-KeyChannel
			net.LookupIP(Sub + "." + *DnslogDomain)
			KeyChannel <- Sub
			i++
			fmt.Println("红队快乐中...，快乐次数：", i)
		}
	}

}
