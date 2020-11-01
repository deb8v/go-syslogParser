package main

import (
	"fmt"
)

/*
0 	Авария (Emergency): система неработоспособна
1 	Тревога (Alert): система требует немедленного вмешательства
2 	Критический (Critical): состояние системы критическое
3 	Ошибка (Error): сообщения о возникших ошибках
4 	Предупреждение (Warning): предупреждения о возможных проблемах
5 	Замечание (Notice): сообщения о нормальных, но важных событиях
6 	Информационный (Informational): информационные сообщения
7 	Отладка (Debug): отладочные сообщения
*/
func getStatusByDict(status int8) string {
	exp := ""

	var m = map[int8]string{
		0: "Emergency",
		1: "Alert",
		2: "Critical",
		3: "Error",
		4: "Warning",
		5: "Notice",
		6: "Informational",
		7: "Debug",
	}

	if status > 0 && status <= 7 {
		//export := (dict[status])
		fmt.Println(m[status])
		//exp = export

	} else {
		exp = "Debug"
	}

	return exp
}
func getRFCByDict(status int8) string {
	exp := ""

	var m = map[int8]string{
		2: "RFC3164 (BSD)",
		1: "RFC5424",
		0: "RAW",
	}

	if status > 0 && status <= 7 {
		//export := (dict[status])
		fmt.Println(m[status])
		//exp = export

	} else {
		exp = "Debug"
	}

	return exp
}
func main() {
	fmt.Println(getStatusByDict(2))
}
