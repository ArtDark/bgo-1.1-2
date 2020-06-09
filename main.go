package main

func main() {
	var balance int32 = 1_500_000_000         // 15 миллионов в копейках
	var transfer int32 = 1_000_000_000        // 10 миллионов в копейках
	total := int64(balance) + int64(transfer) // int32 + int32 будет int32

	println(total)

}
