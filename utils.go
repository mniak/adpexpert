package adpexpert

type coalescible interface {
	string
}

func coalesce[T coalescible](value, fallback T) T {
	var zero T
	if value == zero {
		return fallback
	}
	return value
}
