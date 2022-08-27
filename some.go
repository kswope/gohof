package gohof

func Some[S any](s []S, f func(v S) bool) bool {
	_, found := Find(s, f)
	return found
}
