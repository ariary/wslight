package utils

//Return true if a string is within a slice of string
func Contains(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

//withdraw bytes added in exec combinedOutput (2 bytes added: []byte{13} and \n)
func CleanOutput(output string) string {
	if len(output) > 1 {
		output = output[:len(output)-2] //remove 2 bytes
	}
	return output
}
