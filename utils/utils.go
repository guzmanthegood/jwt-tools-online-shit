package utils

// JoinArray join two slices omitting the duplicity
func JoinArray(a, b []string) []string {
	for _, e := range a {
		if !Contains(b, e) {
			b = append(b, e)
		}
	}
	return b
}

// Contains return true if element are in the list
func Contains(list []string, e string) bool {
	for _, x := range list {
		if x == e {
			return true
		}
	}
	return false
}
