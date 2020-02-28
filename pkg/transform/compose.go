package transform

// StringFn is the basic composable function type
type StringFn func(string) string

// Compose combine two transformers together
func Compose(a StringFn, b ...StringFn) StringFn {
	return func(s string) string {
		if len(b) == 1 {
			return a(b[0](s))
		}
		f := Compose(b[0], b[1:]...)
		return a(f(s))
	}
}
