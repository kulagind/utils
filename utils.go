package utils

func ContainsString(target []string, value string) bool {
	for _, v := range target {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsInt(target []int, value int) bool {
	for _, v := range target {
		if v == value {
			return true
		}
	}
	return false
}
