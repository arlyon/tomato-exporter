package handlers

func quoteme(b []byte) []byte {
	s := []byte("\"")
	b = append(s, b...)
	b = append(b, s...)
	return b
}