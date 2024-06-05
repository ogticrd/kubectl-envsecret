package utils

func RemoveDuplicatedStringE(s []string) []string {
	var cleanedSlice []string
	var set map[string]int = make(map[string]int)

	// Save all items into a set to remove duplicates
	for idx, e := range s {
		set[e] = idx
	}

	// Save all unique items into the cleaned slice
	for k := range set {
		cleanedSlice = append(cleanedSlice, k)
	}

	return cleanedSlice
}
