package service

import (
	"strings"

	"github.com/GraphZC/mq-socket-programming/internal/model"
)

func ExtractCommand(message string) model.Command {
	// Extract command and argument
	commandAndArgs := strings.Split(message, ":")

	args := []string{}

	// Check if it has arguments
	if len(commandAndArgs) > 1 {
		tmp := strings.Split(commandAndArgs[1], ",")
		// Remove whitespace and begin and end of each argument
		for index, arg := range tmp {
			tmp[index] = strings.Trim(arg, " ")
		}
		args = tmp
	}

	return model.Command{
		Cmd:       commandAndArgs[0],
		Arguments: args,
	}
}
