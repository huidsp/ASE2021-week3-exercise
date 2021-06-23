package main

var a2rdata map[int]string

func init() {
	a2rdata = make(map[int]string)
	a2rdata[1] = "I"
	a2rdata[2] = "II"
	a2rdata[3] = "III"
	a2rdata[4] = "IV"
	a2rdata[5] = "V"
	a2rdata[6] = "VI"
	a2rdata[7] = "VII"
	a2rdata[8] = "VIII"
	a2rdata[9] = "IX"
}

func a2r(x int) string {
	if x == 10 {
		return "X"
	} else if x == 20 {
		return "XX"
	}
	return a2rdata[x]
}
