package gohof

func Every[S any](s []S, f func(s S) bool) bool {

	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true

}
