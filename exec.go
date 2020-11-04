package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var argvs struct {
	LockRfc  string `json:"lockRfc"`  //
	Password string `json:"password"` //
	Url      string `json:"url"`      //
	TimeZone string `json:"timeZone"` //
	Host     string `json:"host"`     //
	Port     string `json:"port"`     //
	Address  string `json:"IPAddr"`   //

	//Protocol string `json:"protocol"`   //
	Debug string `json:"debugMsges"` //

}

func makeArgvs(args []string) {
	argvs.LockRfc = "AUTO"
	argvs.Password = ""
	argvs.Address = "0.0.0.0"

	argvs.Debug = "false"
	argvs.Port = "514"

	hn, error := os.Hostname()
	if error != nil {
		hn = "undefined"
	}
	y := string(hn)
	x := strings.SplitN(y, ".", 2)
	argvs.Url = "http://" + x[0] + "/get.go"
	argvs.Host = hn
	for i := range args {
		argsIStr := args[i]
		//fmt.Println(argsIStr)
		argsIStr2 := strings.SplitN(argsIStr, " ", 2)
		fmt.Println(argsIStr2)

		switch argsIStr2[0] {
		case "--lockRfc", "-r":
			argvs.LockRfc = argsIStr2[1]
		case "--password", "-p":
			argvs.Password = argsIStr2[1]
		case "--url", "-u":
			argvs.Url = argsIStr2[1]
		case "--timeZone", "-z":
			argvs.TimeZone = argsIStr2[1]
		case "--host":
			argvs.Host = argsIStr2[1]
		case "--debug":
			argvs.Debug = argsIStr2[1]
		case "--address", "-a":
			argvs.Address = argsIStr2[1]
		case "--port", "-s":
			argvs.Port = argsIStr2[1]
		}

	}
}
func interfacePrint(i interface{}, tableTop string) string {
	s, _ := json.MarshalIndent(i, "", "")
	ss := string(s)
	y := strings.Split(ss, "\n")
	oar := y[1 : len(y)-1]
	outStr := ""
	maxfirstcollen := 10
	maxsecondcollen := 10
	for i := range oar {

		lsi := oar[i]
		lsa := strings.SplitN(lsi, ": ", 2)
		if maxfirstcollen < len(string(lsa[0])) {
			maxfirstcollen = len(string(lsa[0]))
		}
		if maxsecondcollen < len(string(lsa[1])) {
			maxsecondcollen = len(string(lsa[1]))
		}

	}
	tableEnd := "" + strings.Repeat("═", maxfirstcollen+2) + "═" + strings.Repeat("═", maxsecondcollen+2)
	tableDiv := "" + strings.Repeat("─", maxfirstcollen+2) + "┬" + strings.Repeat("─", maxsecondcollen+2)
	tableDivEnd := "" + strings.Repeat("═", maxfirstcollen+2) + "╧" + strings.Repeat("═", maxsecondcollen+2)
	//tableNamingLen := len(tableTop)
	tableNaming := " >>  " + tableTop + strings.Repeat(" ", maxfirstcollen+maxsecondcollen-6)
	//─ ━ │ ┃ ┄ ┅ ┆ ┇ ┈ ┉ ┊ ┋
	//├ ┝ ┞ ┟ ┠ ┡ ┢ ┣ ┤ ┥ ┦ ┧ ┨ ┩ ┪ ┫ ┼ ┽ ┾ ┿ ╀ ╁ ╂ ╃ ╄ ╅ ╆ ╇ ╈ ╉ ╊ ╋
	//┌ ┍ ┎ ┏ ┐ ┑ ┒ ┓ └ ┕ ┖ ┗ ┘ ┙ ┚ ┛
	//┬ ┭ ┮ ┯ ┰ ┱ ┲ ┳ ┴ ┵ ┶ ┷ ┸ ┹ ┺ ┻  ╌ ╍ ╎ ╏ ═ ║
	//╒ ╓ ╔ ╕ ╖ ╗ ╘ ╙ ╚ ╛ ╜ ╝ ╞ ╟ ╠ ╡ ╢ ╣ ╤ ╥ ╦ ╧ ╨ ╩ ╪ ╫ ╬
	outStr += "╔" + tableEnd + "╗\n"
	outStr += "║" + tableNaming + " ║\n"

	outStr += "╟" + tableDiv + "╢\n"
	for i := range oar {

		lsi := oar[i]
		lsa := strings.SplitN(lsi, ": ", 2)
		lsa[1] = strings.TrimRight(lsa[1], ",")
		//outLsi := lsi[0] + string("\t ") + lsi[1] + "\n"
		firstRow := "║ " + lsa[0] + strings.Repeat(" ", (maxfirstcollen+2)-len(lsa[0])-2)
		secondRow := " ┆ " + lsa[1] + strings.Repeat(" ", (maxsecondcollen+2)-len(lsa[1])-2) + " ║\n"
		//outStr += lsa[0] + "\t" + lsa[1] + "\n"
		outStr += firstRow + secondRow

	}
	outStr += "╚" + tableDivEnd + "╝"
	return string(outStr)
}
func parseMessage(str string, thrNum int) {
	fmt.Println(thrNum, str)
	var syslogMeta struct {
		Raw      string `json:"raw"`
		End      string `json:"end"`
		Priority int64  `json:"priority"`
		//PriorityName string `json:"priorityName"`
		Subject     int64  `json:"subject"`
		Topic       string `json:"topics"`
		Timestamp   int64  `json:"timestamp"`
		TimeRFC1123 string `json:"time1123"`
		TimeNow     int64  `json:"timenow"`
		Msg         string `json:"msg"`
		TZ          string `json:"timeZone"` //
		TypeOfName  string `json:"typeOfName"`
	}
	syslogMeta.Raw = str
	priorityRegexp := regexp.MustCompile(`<\d+>`) //5424/3164
	subjectRegexp := regexp.MustCompile(`^\d+ `)  //5424
	strPriorityAfterMask := priorityRegexp.FindString(str)

	if len(strPriorityAfterMask) > 2 {
		priorityIntB, errr := strconv.ParseInt(strPriorityAfterMask[1:len(strPriorityAfterMask)-1], 10, 0)

		if errr != nil {
			panic(errr)
		}
		syslogMeta.Priority = (priorityIntB)
	} else {
		syslogMeta.Priority = 1024
	}
	str = str[len(strPriorityAfterMask):]
	strSubjectAfterMask := subjectRegexp.FindString(str)

	if len(strSubjectAfterMask) >= 2 {
		strSubjectAfterMask = strSubjectAfterMask[0 : len(strSubjectAfterMask)-1]
		subjIntB, errr := strconv.ParseInt(strSubjectAfterMask, 10, 0)

		if errr != nil {
			panic(errr)
		}
		syslogMeta.Subject = (subjIntB)
		syslogMeta.TypeOfName = "RFC5424"
		str = str[len(strSubjectAfterMask)+1:]
	} else {
		syslogMeta.Subject = 1024
		if syslogMeta.Priority != 1024 {
			syslogMeta.TypeOfName = "RFC3164"
		} else {
			syslogMeta.TypeOfName = "RAW"
		}
	}

	//=======================================
	/*
		ANSIC = "Mon Jan _2 15:04:05 2006"
		UnixDate = "Mon Jan _2 15:04:05 MST 2006"
		RubyDate = "Mon Jan 02 15:04:05 -0700 2006"
		RFC822 = "02 Jan 06 15:04 MST"
		RFC822Z = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
		RFC850 = "Monday, 02-Jan-06 15:04:05 MST"
		RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
		RFC1123Z = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
		RFC3339 = "2006-01-02T15:04:05Z07:00"
		RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	*/
	//v1  :1985-04-12T23:20:50.52Z          (ISO 8601 Z-нулевой меридиан) (RFC3339)
	//v2  :1985-04-12T19:20:50.52-04:00
	//v3  :2003-10-11T22:14:15.003Z
	//v4  :2003-08-24T05:14:15.000003-07:00

	//v1.1 :1985-04-12T23:20:50.52Z
	//v2.1 :1985-04-12T19:20:50.52
	//v3.1 :2003-10-11T22:14:15.003Z
	//v4.1 :2003-08-24T05:14:15.000003
	var fromTime time.Time

	timeKnown := false
	var err error = nil
	//localTZ, err := time.LoadLocation(argvs.TimeZone)
	if err != nil {
		fmt.Println("TZ Incorrect")
	}
	//                                      2020-11-04T14:29:20.750334+07:00
	timeRegex5424v21 := regexp.MustCompile(`\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d.\d{6}?[+-]?\d\d:\d\d`)
	timeRegex5424v21Str := timeRegex5424v21.FindString(str)
	if timeRegex5424v21Str != "" {
		fromTime, err = time.Parse("2006-01-02T15:04:05Z07:00", timeRegex5424v21Str)
		timeKnown = true
		fid := timeRegex5424v21.FindStringIndex(timeRegex5424v21Str)

		str = str[fid[1]:len(str)]
	}
	//                                       2020-11-04T14:29:20.750334
	timeRegex5424v211 := regexp.MustCompile(`\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d.\d{6}?Z`)
	timeRegex5424v211Str := timeRegex5424v211.FindString(str)
	if timeRegex5424v211Str != "" {
		fromTime, err = time.Parse("2006-01-02T15:04:05Z", timeRegex5424v211Str)
		timeKnown = true
		fid := timeRegex5424v211.FindStringIndex(timeRegex5424v211Str)

		str = str[fid[1]:len(str)]

	}
	//                                       2020-11-04TT21:29:20.750334Z
	timeRegex5424v31 := regexp.MustCompile(`\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d.\d\d{5}?Z`)
	timeRegex5424v31Str := timeRegex5424v31.FindString(str)
	if timeRegex5424v31Str != "" {
		fromTime, err = time.Parse("2006-01-02T15:04:05Z", timeRegex5424v31Str)
		timeKnown = true
		fid := timeRegex5424v31.FindStringIndex(timeRegex5424v31Str)

		str = str[fid[1]:len(str)]
	}
	// RFC822 (for RFC3164 multiformat parser and damaged RFC5424)
	timeRegex822 := regexp.MustCompile(`\d\d:\d\d:\d\d`)

	if timeKnown == false {
		tzdzStrIndex := timeRegex822.FindStringIndex(str)
		timeRegex822Str := timeRegex822.FindString(str)
		if timeRegex822Str != "" {
			time822Array := strings.Split(timeRegex822Str, ":")
			time822Ih, errr := strconv.ParseInt(time822Array[0], 10, 0)
			time822Im, errr := strconv.ParseInt(time822Array[1], 10, 0)
			time822Is, errr := strconv.ParseInt(time822Array[2], 10, 0)

			if errr != nil {
				panic(errr)
			}
			//add current day to invalid time
			cday := time.Now().Unix() - time.Now().Unix()%86400
			frimInt := (time822Ih * 3600) + (time822Im * 60) + time822Is
			fromTime = time.Unix(cday+frimInt, 0)

			str = str[tzdzStrIndex[1]:len(str)]

		}
	} else {

	}
	if syslogMeta.TypeOfName == "RAW" {
		fromTime = time.Now()

	}
	syslogMeta.Timestamp = fromTime.Unix()
	syslogMeta.TimeRFC1123 = fromTime.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	//timeRegex5424v11 := regexp.MustCompile(`dddd-dd-ddTdd:dd:ddZ07:00`) //2006-01-02T15:04:05Z07:00

	//timeRegex5424v21 := regexp.MustCompile(`dddd-dd-ddTdd:dd:ddZ`)
	//timeRegex5424v21 := regexp.MustCompile(`dddd-dd-ddTdd:dd:ddZ`)
	//timeRegex5424v21 := regexp.MustCompile(`dddd-dd-ddTdd:dd:ddZ`)
	//fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05Z07:00"))
	//t2, err := time.Parse("2006-01-02T15:04:05Z07:00")
	if err != nil {
		panic(err)
	} //RFC1123Z

	//timeRegex5424v4 := regexp.MustCompile(`dddd-dd-ddTdd:dd:dd`)
	syslogMeta.TZ = argvs.TimeZone
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")

	//5424 examples SDPARAM
	/*
			1.
			[exampleSDID@32473 iut="3" eventSource="Application"
			eventID="1011"]

			2.
			[exampleSDID@32473 iut="3" eventSource="Application"
		 	eventID="1011"][examplePriority@32473 class="high"]
	*/
	if syslogMeta.TypeOfName == "RFC3164" {

	}
	syslogMeta.End = str
	if argvs.Debug == "true" || argvs.Debug == "TRUE" {
		fmt.Println(interfacePrint(syslogMeta, string(syslogMeta.Raw)))
	}
	//fmt.Println("end")

}

func listener(messages chan string) {
	for { //в вечном цикле ждём что появилось нового
		//parseMessage(<-messages, rand.Intn(1024)) //запуск аснхронного потока
		select {
		case msg := <-messages: //получаем сообщение в обработку повешав цикл
			parseMessage(msg, rand.Intn(1024)) //запуск аснхронного потока
		}
	}

}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	//argsArray := strings.Split(argsWithoutProg[0], " ")
	makeArgvs(argsWithoutProg)

	fmt.Println(argsWithProg)
	fmt.Println(interfacePrint(argvs, "argvs"))
	messages := make(chan string)
	go listener(messages)
	if false {

		conn, err := net.ListenUDP("udp", &net.UDPAddr{
			Port: 514,
			IP:   net.ParseIP("0.0.0.0"),
		})
		if err != nil {
			panic(err)
		}

		defer conn.Close()
		fmt.Printf("server listening %s\n", conn.LocalAddr().String())

		for {
			message := make([]byte, 4096)
			rlen, remote, err := conn.ReadFromUDP(message[:])
			if err != nil {
				panic(err)
			}

			data := strings.TrimSpace(string(message[:rlen]))
			messages <- data
			if argvs.Debug == "true" || argvs.Debug == "TRUE" {
				fmt.Printf("from %s \n", remote)
			}
		}

	}
	//5424
	//v1  :1985-04-12T23:20:50.52Z          (ISO 8601 Z-нулевой меридиан) (RFC3339)
	//v2  :1985-04-12T19:20:50.52-04:00
	//v3  :2003-10-11T22:14:15.003Z
	//v4  :2003-08-24T05:14:15.000003-07:00
	examples := [...]string{ // Компилятор Go подсчитывает элементы
		`<14>1 2020-11-04T14:29:20.750334+07:00 DESKTOP-PST4DMA t3 - - [timeQuality tzKnown="1" isSynced="0"] ebal<22>;TABLE DROM TEST; tzv2`,
		`<14>1 2020-11-04T14:29:20.750334Z DESKTOP-PST4DMA t3 - - [timeQuality tzKnown="1" isSynced="0"] ebal<22>;TABLE DROM TEST; tzv2.1`,
		`<14>1 2020-11-04T21:29:20.750334Z DESKTOP-PST4DMA t3 - - [timeQuality tzKnown="1" isSynced="0"] ebal<22>;TABLE DROM TEST; tzv1`,
		`<14>1 2020-11-04T21:29:20.750334Z DESKTOP-PST4DMA t3 - - [timeQuality tzKnown="1" isSynced="0"] ebal<22>;TABLE DROM TEST; tzv3`,
		`<14>1 2020-11-04T14:29:20.750334+07:00 DESKTOP-PST4DMA t3 - - [timeQuality tzKnown="1" isSynced="0"] ebal<22>;TABLE DROM TEST; tzv4`,
		`<12>Nov  4 14:29:32 DESKTOP-PST4DMA t2: ebal`,
		//mikrotik log (bsd 16,0)
		`<128>Nov  4 14:30:03 MikroTik 0.18b, info: user admin logged in from 192.168.88.238 via telnet`,
		`<134>Nov  4 14:31:36 MikroTik 0.18b, info: log action changed by admin`,
		`system,info 0.18b, info: log action changed by admin`,
		`Сатурн`,
		//`alarm Сатурн`,
	}
	if false {

		time.Sleep(16000 * time.Millisecond)
	}
	//parseMessage(msg, rand.Intn(1024)) //запуск аснхронного потока
	for i := range examples {
		parseMessage(examples[i], rand.Intn(1024)) //запуск аснхронного потока
	}
}
