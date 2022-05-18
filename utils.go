package adpexpert

func coalesce(value, fallback string) string {
	var zero string
	if value == zero {
		return fallback
	}
	return value
}
