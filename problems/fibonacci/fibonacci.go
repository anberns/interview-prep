package fibonacci

func FibIter(num int) []int {
	seq := make([]int, num)

	if num == 0 || num == 1 {
		return seq
	} else if num == 2 {
		seq[1] = 1
		return seq
	} else {
		fibIter(num, seq)
		return seq
	}
}

func fibIter(num int, seq []int) {
	first := 0
	second := 1

	seq[0] = first
	seq[1] = second

	for i := 2; i < num; i++ {
		val := first + second
		seq[i] = val
		first = second
		second = val
	}
}

func FibRec(n int) int {

	mem := make(map[int]int, n)
	mem[0] = 0
	mem[1] = 1

	return fibRec(n, mem)
}

func fibRec(n int, mem map[int]int) int {
	if n < 2 {
		return n
	}

	var val1 int
	var val2 int

	v1, ok := mem[n-1]
	if ok {
		val1 = v1
	} else {
		val1 = fibRec(n-1, mem)
	}

	v2, ok := mem[n-2]
	if ok {
		val2 = v2
	} else {
		val2 = fibRec(n-2, mem)
	}

	return val1 + val2
}
