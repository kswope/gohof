package gohof

func Map[S any, T any](s []S, f func(S) T) []T {
	var accum []T
	for _, v := range s {
		accum = append(accum, f(v))
	}
	return accum
}
