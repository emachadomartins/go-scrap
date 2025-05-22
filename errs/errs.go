package errs

func TryCatch[T any](
	value T,
	err error,
) T {
	if err != nil {
		panic(err)
	}
	return value
}

func ThrowOnError(err error) {
	if err != nil {
		panic(err)
	}
}
