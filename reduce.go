package gohof

func Reduce[S any, T any](s []S, f func(a T, v S) T, initial T) T {
	accum := initial
	for _, v := range s {
		accum = f(accum, v)
	}
	return accum
}
