package color

import "strings"

func Colorize(log string, key string) string {
	switch strings.ToUpper(key) {
	case "DEBUG":
		return "\033[38;2;170;170;180m" + log + "\033[0m"
	case "INFO":
		return "\033[38;2;100;235;145m" + log + "\033[0m"
	case "WARN":
		return "\033[38;2;255;210;90m" + log + "\033[0m"
	case "ERROR":
		return "\033[38;2;255;110;120m" + log + "\033[0m"
	case "FATAL":
		return "\033[1;38;2;255;80;100m" + log + "\033[0m"
	case "TIME", "DURATION":
		return "\033[38;2;145;190;255m" + log + "\033[0m"
	case "MSG":
		return "\033[38;2;240;240;250m" + log + "\033[0m"
	case "REQUESTID":
		return "\033[38;2;210;170;255m" + log + "\033[0m"
	case "METHOD":
		return "\033[38;2;110;230;205m" + log + "\033[0m"
	case "PATH":
		return "\033[38;2;180;220;145m" + log + "\033[0m"
	case "STATUS":
		return "\033[38;2;225;170;255m" + log + "\033[0m"
	default:
		return "\033[38;2;220;220;235m" + log + "\033[0m"
	}
}
