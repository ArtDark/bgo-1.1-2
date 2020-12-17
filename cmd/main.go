package main

func main() {
	var balance int64 = 1_500_000_000  // 15 миллионов в копейках
	var transfer int64 = 1_000_000_000 // 10 миллионов в копейках
	total := balance + transfer

	println(total)

}
