package color

import "strings"

func Colorize(log string, key string) string {
	switch strings.ToUpper(key) {
	case "DEBUG":
		return "\033[1;38;2;190;190;200m" + log + "\033[0m"
	case "INFO":
		return "\033[1;38;2;60;255;120m" + log + "\033[0m"
	case "WARN":
		return "\033[1;38;2;255;225;40m" + log + "\033[0m"
	case "ERROR":
		return "\033[1;38;2;255;75;85m" + log + "\033[0m"
	case "FATAL":
		return "\033[1;38;2;255;35;70m" + log + "\033[0m"
	case "TIME", "DURATION":
		return "\033[1;38;2;100;185;255m" + log + "\033[0m"
	case "MSG":
		return "\033[1;38;2;255;255;255m" + log + "\033[0m"
	case "REQUESTID":
		return "\033[1;38;2;210;135;255m" + log + "\033[0m"
	case "METHOD":
		return "\033[1;38;2;50;255;205m" + log + "\033[0m"
	case "PATH":
		return "\033[1;38;2;180;255;100m" + log + "\033[0m"
	case "STATUS":
		return "\033[1;38;2;235;125;255m" + log + "\033[0m"
	default:
		return "\033[1;38;2;245;245;255m" + log + "\033[0m"
	}
}
