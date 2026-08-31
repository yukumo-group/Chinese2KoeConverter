package process

// ConvertBoolToString converts bool val to string val
func ConvertBoolToString(
	data bool,
) string {
	if data {
		return "true"
	}
	return "false"
}
