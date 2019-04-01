package util

func WrapPage(page int, size int, length int) (int, int) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 1
	}

	start := (page - 1) * size

	if length < start {
		return 0, 0
	}

	surplus := length - start
	if surplus < size {
		size = surplus
	}

	return start , start+size
}
