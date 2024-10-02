package main

func main() {
	var (
		a, b = 10, 20
	)
	println("sum:", (a + b))
	{
		var c float32 = 21.23
		var ok = true
		{
			var d = 12.23
			println("c,ok,d", c, ok, d)
		}

		// println(d)
	}
	var e = 123.42
	println("e", e)
}
