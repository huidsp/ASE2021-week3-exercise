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

	a2rdata[10] = "X"
	a2rdata[20] = "XX"
	a2rdata[30] = "XXX"
	a2rdata[40] = "XL"
	a2rdata[50] = "L"
	a2rdata[60] = "LX"
	a2rdata[70] = "LXX"
	a2rdata[80] = "LXXX"
	a2rdata[90] = "XC"
}

func a2r(x int) string {
	if x >= 10 {
		if x == 11 {
			return a2rdata[10] + a2rdata[1]
		}
		return a2rdata[x]
	}
	return a2rdata[x]
}
