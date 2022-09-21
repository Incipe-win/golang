package main

type T struct {
	a int
}

func (t T) M1() {
}

func (t *T) M2() {
	t.a = 11
}

func main() {
	var t T
	println(t.a)
	t.M1()
	println(t.a)
	t.M2()
	println(t.a)
}
