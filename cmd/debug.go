package cmd

import "fmt"

func Debugf(format string, a ...interface{}) {
	if debug {
		fmt.Printf(format, a)
	}
}

func Debugln(a ...interface{}) {
	if debug {
		fmt.Println(a)
	}
}

func Debugv(a interface{}) {
	if debug {
		fmt.Printf("%+v\n", a)
	}
}
