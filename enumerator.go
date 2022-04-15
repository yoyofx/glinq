package glinq

func DoWhile[T any](enumerator IEnumerator[T], action func(T) bool) {
	for {
		if v, ok := enumerator.Value(); ok {
			if !action(v) {
				break
			}
		}
		if !enumerator.Next() {
			break
		}
	}
	enumerator.Reset()
}
