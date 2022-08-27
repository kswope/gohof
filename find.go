package gohof

func Find[S any](s []S, f func(S) bool) (S, bool) {

	for _, v := range s {
		if f(v) {
			return v, true
		}
	}

	var zeroValue S
	return zeroValue, false

}
