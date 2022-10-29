package utils

func Contains(target []string, value string) bool {
	for _, v := range target {
		if v == value {
			return true
		}
	}
	return false
}
