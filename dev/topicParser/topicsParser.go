package main

import (
	"fmt"
	"strings"
)

//typeOf 1
func getTopicsfromRAW(raw string) string {
	e := strings.Split(raw, ":")[0]
	e = strings.ReplaceAll(e, " ", ",")
	a := strings.Split(e, ",")
	topics := ""
	for index, element := range a {
		if index < len(a)-2 {
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
	stamp3164 := "<12>Nov  1 20:57:13 DESKTOP-DDDAAA t2: ebal"
	stamp5424 := "<12>1 2020-11-01T20:47:57.389094+07:00 DESKTOP-DDDAAA t2 - - [timeQuality tzKnown=\"1\" isSynced=\"0\"] ebal"

	fmt.Printf("export:  %s\n", stamp5424)
	fmt.Printf("export:  %s\n", getTopicsfromRFC5424(stamp5424))
	fmt.Printf("export:  %s\n", stamp3164)
	fmt.Printf("export:  %s\n", getTopicsfromRFC3164(stamp3164))
	fmt.Printf("export:  %s\n", stampRAW)
	fmt.Printf("export:  %s\n", getTopicsfromRAW(stampRAW))
}
