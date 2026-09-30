package formatter

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/Weit145/simple-log/internal/color"
)

func FormatLog(log map[string]interface{}) string {
	keys := make([]string, 0, len(log))
	for key := range log {
		keys = append(keys, key)
	}

	priority := map[string]int{"time": 0, "level": 1, "msg": 2}
	sort.Slice(keys, func(i, j int) bool {
		left, leftKnown := priority[keys[i]]
		right, rightKnown := priority[keys[j]]
		if !leftKnown {
			left = 3
		}
		if !rightKnown {
			right = 3
		}
		if left != right {
			return left < right
		}
		return keys[i] < keys[j]
	})

	var result strings.Builder
	for _, key := range keys {
		if result.Len() > 0 {
			result.WriteByte('\t')
		}
		preparedValue := key + ": "
		if value, ok := log[key].(string); ok {
			preparedValue += value
		} else {
			encoded, _ := json.Marshal(log[key])
			preparedValue += string(encoded)
		}
		colorKey := key
		if key == "level" {
			if level, ok := log[key].(string); ok {
				colorKey = level
			}
		}
		result.WriteString(color.Colorize(preparedValue, colorKey))
	}
	return result.String()
}
