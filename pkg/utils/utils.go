package utils

func Contains(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}
