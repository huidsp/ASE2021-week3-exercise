package main

var r2adata map[string]int

func init() {
	r2adata = make(map[string]int)
	r2adata["I"] = 1
	r2adata["II"] = 2
}

func r2a(x string) int {
	return r2adata[x]
}
