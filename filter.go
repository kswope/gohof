package gohof

func Filter[S any](s []S, f func(S) bool) []S {
	var accum []S
	for _, v := range s {
		if f(v) {
			accum = append(accum, v)
		}
	}
	return accum
}
