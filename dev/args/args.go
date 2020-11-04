package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
		if maxfirstcollen < len(lsa[0]) {
			maxfirstcollen = len(lsa[0])
		}
		if maxsecondcollen < len(lsa[1]) {
			maxsecondcollen = len(lsa[1])
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
func main() {

	// `os.Args` позволяет получить доступ к аргументам
	// командной строки. Обратите внимание, что первое
	// значение это путь к самой программе, а
	// `os.Args[1:]` содержит только аргументы.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	//argsArray := strings.Split(argsWithoutProg[0], " ")
	makeArgvs(argsWithoutProg)

	fmt.Println(argsWithProg)
	if argvs.Debug == "true" || argvs.Debug == "TRUE" {
		fmt.Println(interfacePrint(argvs, "argvs"))
	}

	//interfacePrint(argvs, "argvs")

}
