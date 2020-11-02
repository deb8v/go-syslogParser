package main

import (
	"fmt"
	"strings"
)

//typeOf 1
func getMsgfromRAW(raw string) string {
	e := strings.SplitAfterN(raw, ":", 2)
	msg := ""
	if len(e) < 2 {
		e = strings.SplitAfterN(raw, "\t", 2)
	} else {
		if len(e) < 2 {
			e = strings.SplitAfterN(raw, "on", 2)
		} else {
			if len(e) < 2 {
				e = strings.SplitAfterN(raw, "     ", 2)
			}
		}
	}
	if len(e) < 2 {
		msg = e[0]
	} else {
		msg = e[1]
	}
	return msg
}
func getMsgfromRFC5424(raw string) string {
	e := strings.Split(raw, "- ")
	g := e[2]
	export := ""
	if strings.Count(g, "]") == 1 {
		export = strings.Split(g, "]")[1]
	}
	if strings.Count(g, "]") == 0 {
		export = strings.Split(raw, "- ")[2]
	}

	return strings.TrimSpace(export)

}

func getMsgfromRFC3164(raw string) string {
	e := strings.Split(raw, ":")

	return strings.TrimSpace(e[3])
}
func main() {
	stampRAW := "system,info DESKTOP-AAAAAA: ebal"
	stampRAW_2 := "dns,packet <tracker.tiny-vps.com:CNAME:422=tardis.tiny-vps.com>"
	stampRAW_3 := "dhcp,debug,packet     Parameter-List = Subnet-Mask,Broadcast-Address,Unknown(2),Router,Domain-Name,Domain-Server,Domain-Search,Host-Name,NETBIOS-Name-Server,NETBIOS-Scope,Interface-MTU,Classless-Route,NTP-Server"
	stamp3164 := "<12>Nov  1 20:57:13 DESKTOP-DDDAAA t2: ebal"
	stamp5424 := "<12>1 2020-11-01T20:47:57.389094+07:00 DESKTOP-DDDAAA t2 - - [timeQuality tzKnown=\"1\" isSynced=\"0\"] ebal"
	stamp5424_2 := "<12>1 2020-11-01T20:47:57.389094+07:00 DESKTOP-DDDAAA t2 - - ebal"
	fmt.Printf("5424:       \t%s\n", stamp5424)
	fmt.Printf("5424:       \t%s\n", getMsgfromRFC5424(stamp5424))
	fmt.Printf("5424_2:     \t%s\n", stamp5424_2)
	fmt.Printf("5424_2:     \t%s\n", getMsgfromRFC5424(stamp5424_2))
	fmt.Printf("3164:       \t%s\n", stamp3164)
	fmt.Printf("3164:       \t%s\n", getMsgfromRFC3164(stamp3164))
	fmt.Printf("stampRAW1:  \t%s\n", stampRAW)
	fmt.Printf("stampRAW1:  \t%s\n", getMsgfromRAW(stampRAW))
	fmt.Printf("stampRAW2:  \t%s\n", stampRAW_2)
	fmt.Printf("stampRAW2:  \t%s\n", getMsgfromRAW(stampRAW_2))
	fmt.Printf("stampRAW3:  \t%s\n", stampRAW_3)
	fmt.Printf("stampRAW3:  \t%s\n", getMsgfromRAW(stampRAW_3))
}
