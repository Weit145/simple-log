package color

import "strings"

func Colorize(log string, key string) string {
	switch strings.ToUpper(key) {
	case "DEBUG":
		return "\033[90m" + log + "\033[0m"
	case "INFO":
		return "\033[38;2;120;180;140m" + log + "\033[0m"
	case "WARN":
		return "\033[38;2;210;170;100m" + log + "\033[0m"
	case "ERROR":
		return "\033[38;2;210;115;115m" + log + "\033[0m"
	case "FATAL":
		return "\033[38;2;190;90;100m" + log + "\033[0m"
	case "TIME", "DURATION":
		return "\033[38;2;140;160;175m" + log + "\033[0m"
	case "MSG":
		return "\033[38;2;180;180;185m" + log + "\033[0m"
	case "REQUESTID":
		return "\033[38;2;150;145;175m" + log + "\033[0m"
	case "METHOD":
		return "\033[38;2;130;175;160m" + log + "\033[0m"
	case "PATH":
		return "\033[38;2;155;170;150m" + log + "\033[0m"
	case "STATUS":
		return "\033[38;2;170;150;180m" + log + "\033[0m"
	default:
		return "\033[38;2;160;160;170m" + log + "\033[0m"
	}
}
