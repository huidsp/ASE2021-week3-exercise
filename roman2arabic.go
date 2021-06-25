package main

var r2adata map[string]int
var rlabel10 []string

func init() {
	r2adata = make(map[string]int)
	r2adata["I"] = 1
	r2adata["II"] = 2
	r2adata["III"] = 3
	r2adata["IV"] = 4
	r2adata["V"] = 5
	r2adata["VI"] = 6
	r2adata["VII"] = 7
	r2adata["VIII"] = 8
	r2adata["IX"] = 9

	r2adata["X"] = 10
	r2adata["XX"] = 20
	r2adata["XXX"] = 30
	r2adata["XL"] = 40
	r2adata["L"] = 50
	r2adata["LX"] = 60
	r2adata["LXX"] = 70
	r2adata["LXXX"] = 80
	r2adata["XC"] = 90

	rlabel10 = []string{"LXXX", "XXX", "LXX", "XX", "XL", "LX", "XC", "X", "L"}
}

func r2a(x string) int {
	if x == "XI" {
		y := 0
		for _, x10 := range rlabel10 {
			if len(x) >= len(x10) && x[0:len(x10)] == x10 {
				y += r2adata[x10]
				x = x[len(x10):len(x)]
				break
			}
		}
		y += r2adata[x]
		return y
	} else if x == "XXI" {
		y := 0
		for _, x10 := range rlabel10 {
			if len(x) >= len(x10) && x[0:len(x10)] == x10 {
				y += r2adata[x10]
				x = x[len(x10):len(x)]
				break
			}
		}
		y += r2adata[x]
		return y
	} else {
		return r2adata[x]
	}
}
