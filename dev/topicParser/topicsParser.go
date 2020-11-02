package main

import (
	"fmt"
	"strings"
)

//typeOf 1
func getTopicsfromRAW(raw string) string {
	c := strings.SplitAfterN(raw, ":", 2)
	if len(c) < 2 {
		c = strings.SplitAfterN(raw, "\t", 2)
	}
	if len(c) < 2 {
		c = strings.SplitAfterN(raw, "on", 2)
	}
	if len(c) < 2 {
		c = strings.SplitAfterN(raw, "     ", 2)
	}
	e := strings.TrimSpace(c[0])
	e = strings.ReplaceAll(e, " ", ",")
	a := strings.Split(e, ",")
	topics := ""
	for index, element := range a {
		if index < len(a)-1 {
			topics += element + ", "
		} else {
			topics += element
		}

	}

	return topics
}
func getTopicsfromRFC5424(raw string) string {
	e := strings.Split(raw, " ")

	topics := ""
	if len(e) > 4 {
		topics = e[2] + ", " + e[3]
	}
	return topics
}

func getTopicsfromRFC3164(raw string) string {
	e := strings.Split(raw, " ")

	topics := ""
	if len(e) > 6 {
		topics = e[4] + ", " + e[5]
	}
	return topics
}
func main() {
	stampRAW := "system,info DESKTOP-AAAAAA: ebal"
	stampRAW_2 := "dns,packet <tracker.tiny-vps.com:CNAME:422=tardis.tiny-vps.com>"
	stampRAW_3 := "dhcp,debug,packet     Parameter-List = Subnet-Mask,Broadcast-Address,Unknown(2),Router,Domain-Name,Domain-Server,Domain-Search,Host-Name,NETBIOS-Name-Server,NETBIOS-Scope,Interface-MTU,Classless-Route,NTP-Server"
	stamp3164 := "<12>Nov  1 20:57:13 DESKTOP-DDDAAA t2: ebal"
	stamp5424 := "<12>1 2020-11-01T20:47:57.389094+07:00 DESKTOP-DDDAAA t2 - - [timeQuality tzKnown=\"1\" isSynced=\"0\"] ebal"

	fmt.Printf("export:  %s\n", stamp5424)
	fmt.Printf("export:  %s\n", getTopicsfromRFC5424(stamp5424))
	fmt.Printf("export:  %s\n", stamp3164)
	fmt.Printf("export:  %s\n", getTopicsfromRFC3164(stamp3164))
	fmt.Printf("export:  %s\n", stampRAW)
	fmt.Printf("export:  %s\n", getTopicsfromRAW(stampRAW))
	fmt.Printf("export:  %s\n", stampRAW_2)
	fmt.Printf("export:  %s\n", getTopicsfromRAW(stampRAW_2))
	fmt.Printf("export:  %s\n", stampRAW_3)
	fmt.Printf("export:  %s\n", getTopicsfromRAW(stampRAW_3))
}
