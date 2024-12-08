package balanced

func EqualFrequency(word string) bool {
	if len(word) <= 3 {
		return true
	}

	m1 := make(map[string]int)
	for i := 0; i < len(word); i++ {
		str := string(word[i])
		m1[str] += 1
	}

	m2 := make(map[int]int)
	for _, v := range m1 {
		m2[v] += 1
	}

	if len(m2) <= 1 {
		for k, v := range m2 {
			return k == 1 || v == 1
		}
	}

	if len(m2) > 2 {
		return false
	}

	if len(m2) == 2 {
		var k1, k2, v1, v2 int
		i := 1
		for k, v := range m2 {
			if i == 1 {
				k1 = k
				v1 = v
				i++
			} else if i == 2 {
				k2 = k
				v2 = v
				i++
			}
		}
		if v1 == 1 && v2 == 1 {
			return (k1-1 == 0 || k1-1 == k2) || (k2-1 == 0 || k2-1 == k1)
		} else if v1 == 1 {
			return k1-1 == 0 || k1-1 == k2
		} else if v2 == 1 {
			return k2-1 == 0 || k2-1 == k1
		}
	}

	return false
}
