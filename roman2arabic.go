package main

var r2adata map[string]int

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
}

func r2a(x string) int {
	return r2adata[x]
}
