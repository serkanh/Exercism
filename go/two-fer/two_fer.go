package twofer

// ShareWith takes a string or null value and returns a string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
