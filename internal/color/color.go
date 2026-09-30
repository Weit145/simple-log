package color

import "strings"

func Colorize(log string, key string) string {
	switch strings.ToUpper(key) {
	case "DEBUG":
		return "\033[90m" + log + "\033[0m"
	case "INFO":
		return "\033[32m" + log + "\033[0m"
	case "WARN":
		return "\033[33m" + log + "\033[0m"
	case "ERROR":
		return "\033[31m" + log + "\033[0m"
	case "FATAL":
		return "\033[1;31m" + log + "\033[0m"
	case "TIME":
		return "\033[36m" + log + "\033[0m"
	default:
		return "\033[35m" + log + "\033[0m"
	}
}
