package main

import (
	"fmt"
	"strings"
)

//typeOf 1
func getMsgfromRAW(raw string) string {
	e := strings.Split(raw, ": ")[1]

	return e
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
	stamp3164 := "<12>Nov  1 20:57:13 DESKTOP-DDDAAA t2: ebal"
	stamp5424 := "<12>1 2020-11-01T20:47:57.389094+07:00 DESKTOP-DDDAAA t2 - - [timeQuality tzKnown=\"1\" isSynced=\"0\"] ebal"
	stamp5424_2 := "<12>1 2020-11-01T20:47:57.389094+07:00 DESKTOP-DDDAAA t2 - - ebal"
	fmt.Printf("export:  %s\n", stamp5424)
	fmt.Printf("export:  %s\n", getMsgfromRFC5424(stamp5424))
	fmt.Printf("export:  %s\n", stamp5424_2)
	fmt.Printf("export:  %s\n", getMsgfromRFC5424(stamp5424_2))
	fmt.Printf("export:  %s\n", stamp3164)
	fmt.Printf("export:  %s\n", getMsgfromRFC3164(stamp3164))
	fmt.Printf("export:  %s\n", stampRAW)
	fmt.Printf("export:  %s\n", getMsgfromRAW(stampRAW))
}
