package grid

func FindAllDuplicates(slices [][]int) []int {
	result := slices[0]

	for i := 1; i < len(slices); i++ {
		result = FindDuplicates(result, slices[i])
	}

	return result
}

func FindDuplicates(slice1 []int, slice2 []int) []int {
	var src []int
	var tgt []int
	var result []int

	if len(slice1) < len(slice2) {
		src = slice1
		tgt = slice2
	} else {
		src = slice2
		tgt = slice1
	}

	for _, v := range src {
		if Contains(tgt, v) && !Contains(result, v) {
			result = append(result, v)
		}
	}

	return result
}

func FindUnique(slice []int) []int {
	unique := []int{}
	m := map[int]bool{}

	for _, v := range slice {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func Contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
