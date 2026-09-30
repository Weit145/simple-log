package parse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Weit145/simple-log/internal/formatter"
)

func Parse(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Bytes()

		var result map[string]interface{}
		if err := json.Unmarshal(line, &result); err != nil {
			if _, err := fmt.Fprintln(output, string(line)); err != nil {
				return err
			}
			continue
		}

		if _, err := fmt.Fprintln(output, formatter.FormatLog(result)); err != nil {
			return err
		}

	}

	return scanner.Err()
}
