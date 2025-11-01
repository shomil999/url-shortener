package shortener

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func base62Encode(b []byte) string {
	if len(b) == 0 {
		return "0"
	}
	n := make([]byte, len(b))
	copy(n, b)
	var i int
	for i < len(n) && n[i] == 0 {
		i++
	}
	n = n[i:]
	if len(n) == 0 {
		return "0"
	}
	var out []byte
	for len(n) > 0 {
		quot := make([]byte, 0, len(n))
		rem := 0
		for _, v := range n {
			cur := int(v) + rem*256
			q := cur / 62
			rem = cur % 62
			if len(quot) > 0 || q > 0 {
				quot = append(quot, byte(q))
			}
		}
		out = append(out, alphabet[rem])
		n = quot
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}