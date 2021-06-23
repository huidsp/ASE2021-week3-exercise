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

	a2rdata[100] = "C"
	a2rdata[200] = "CC"
	a2rdata[300] = "CCC"
	a2rdata[400] = "CD"
	a2rdata[500] = "D"
	a2rdata[600] = "DC"
	a2rdata[700] = "DCC"
	a2rdata[800] = "DCCC"
	a2rdata[900] = "CM"

	a2rdata[1000] = "M"

}

func a2r(x int) string {
	str := ""
	if x >= 1000 {
		return a2rdata[x]
	}
	if x >= 100 {
		hundred := (x / 100) * 100
		x -= hundred
		str += a2rdata[hundred]
	}
	if x >= 10 {
		ten := (x / 10) * 10
		x -= ten
		str += a2rdata[ten]
	}
	str += a2rdata[x]
	return str
}
