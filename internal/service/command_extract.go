package service

import (
	"strings"

	"github.com/GraphZC/mq-socket-programming/internal/model"
)

func ExtractCommand(message string) model.Command {
	commandAndArgs := strings.Split(message, ":")

	args := []string{}
	if len(commandAndArgs) > 1 {
		tmp := strings.Split(commandAndArgs[1], ",")
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
