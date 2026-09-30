package parse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Weit145/simple-log/internal/formatter"
)

func Parse() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Bytes()

		var result map[string]interface{}
		if err := json.Unmarshal(line, &result); err != nil {
			fmt.Println(string(line))
			continue
		}

		fmt.Println(formatter.FormatLog(result))

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
	}
}
