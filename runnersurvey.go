// +build linux darwin windows openbsd netbsd freebsd

package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func RunnerGroupSurvey(taskAgentPool string, taskAgentPools []string) string {
	prompt := &survey.Select{
		Message: "Choose a runner group:",
		Options: taskAgentPools,
	}
	err := survey.AskOne(prompt, &taskAgentPool)
	if err != nil {
		fmt.Println("Failed to retrieve your choice using default runner group: " + taskAgentPool)
	}
	return taskAgentPool
}

func GetInput(prompt string, answer string) string {
	in := &survey.Input{
		Message: prompt,
		Default: answer,
	}
	if err := survey.AskOne(in, &answer); err != nil {
		fmt.Println("Failed to retrieve your choice using default: " + answer)
	}
	return answer
}
