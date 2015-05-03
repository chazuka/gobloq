package sort

const (
	ASC  = 1
	DESC = 2
)

func Int(a []int, d int) []int {

	for i := len(a) - 1; i >= 0; i-- {
		current := a[i]
		position := i
		for j := i - 1; j >= 0; j-- {
			prev := a[j]
			if (d == ASC && current < prev) || (d == DESC && current > prev) {
				current = prev
				position = j
			}
		}

		if position != i {
			t := a[position]
			a[position] = a[i]
			a[i] = t
		}
	}

	return a
}

func IntAsc(a []int) []int {
	return Int(a, ASC)
}

func IntDesc(a []int) []int {
	return Int(a, DESC)
}
