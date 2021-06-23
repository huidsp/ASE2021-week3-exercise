package main

var a2rdata map[int]string

func init() {
	a2rdata = make(map[int]string)
	a2rdata[1] = "I"
	a2rdata[2] = "II"
	a2rdata[3] = "III"
}

func a2r(x int) string {
	return a2rdata[x]
}
