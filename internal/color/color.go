package color

import "strings"

func Colorize(log string, key string) string {
	switch strings.ToUpper(key) {
	case "DEBUG":
		return "\033[1;37m" + log + "\033[0m"
	case "INFO":
		return "\033[1;92m" + log + "\033[0m"
	case "WARN":
		return "\033[1;93m" + log + "\033[0m"
	case "ERROR":
		return "\033[1;91m" + log + "\033[0m"
	case "FATAL":
		return "\033[1;91m" + log + "\033[0m"
	case "TIME", "DURATION":
		return "\033[1;96m" + log + "\033[0m"
	case "MSG":
		return "\033[1;97m" + log + "\033[0m"
	case "REQUESTID":
		return "\033[1;95m" + log + "\033[0m"
	case "METHOD":
		return "\033[1;96m" + log + "\033[0m"
	case "PATH":
		return "\033[1;94m" + log + "\033[0m"
	case "STATUS":
		return "\033[1;93m" + log + "\033[0m"
	default:
		return "\033[1;97m" + log + "\033[0m"
	}
}
