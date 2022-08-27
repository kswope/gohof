package gohof

func Partial1[S any, T any](f func(x S, y S) T, x S) func(S) T {
	return func(y S) T {
		return f(x, y)
	}
}
