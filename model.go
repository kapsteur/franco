package franco

type Tuple struct {
	Code  string
	Count float64
}

type Tuples []Tuple

func (t Tuples) Len() int {
	return len(t)
}

func (t Tuples) Less(i, j int) bool {
	return t[i].Count < t[j].Count
}

func (t Tuples) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type Languages map[string]Trigrams
type Trigrams map[string]int
