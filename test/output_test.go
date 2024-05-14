package test

import (
	"fmt"
	"testing"
	"ascii/methods"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
)

type tests struct {
	name     string
	input    string
	expected string
}

func TestProccessOutput(t *testing.T) {
	graphics := methods.ProccessFileContent("../standard.txt")

	cases := []tests{
		{
			name:     "Hello",
			input:    "Hello",
			expected: " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n",
		},
		{
			name:     `1 Hello\n2 World`,
			input:    `1 Hello\n2 World`,
			expected: "           _    _          _   _          \n _        | |  | |        | | | |         \n/ |       | |__| |   ___  | | | |   ___   \n| |       |  __  |  / _ \\ | | | |  / _ \\  \n| |       | |  | | |  __/ | | | | | (_) | \n|_|       |_|  |_|  \\___| |_| |_|  \\___/  \n                                          \n                                          \n              __          __                 _       _  \n ____         \\ \\        / /                | |     | | \n|___ \\         \\ \\  /\\  / /    ___    _ __  | |   __| | \n  __) |         \\ \\/  \\/ /    / _ \\  | '__| | |  / _` | \n / __/           \\  /\\  /    | (_) | | |    | | | (_| | \n|_____|           \\/  \\/      \\___/  |_|    |_|  \\__,_| \n                                                        \n                                                        \n",
		},
		{
			name:     `new\n\nlines`,
			input:    `new\n\nlines`,
			expected: "                          \n                          \n _ __     ___  __      __ \n| '_ \\   / _ \\ \\ \\ /\\ / / \n| | | | |  __/  \\ V  V /  \n|_| |_|  \\___|   \\_/\\_/   \n                          \n                          \n\n _   _                       \n| | (_)                      \n| |  _   _ __     ___   ___  \n| | | | | '_ \\   / _ \\ / __| \n| | | | | | | | |  __/ \\__ \\ \n|_| |_| |_| |_|  \\___| |___/ \n                             \n                             \n",
		},
		{
			name:     `Don't show \r\f\v\b\a`,
			input:    "Don't show \r\f\v\b\a",
			expected: " _____                    _   _                 _                               \n|  __ \\                  ( ) | |               | |                              \n| |  | |   ___    _ __   |/  | |_         ___  | |__     ___   __      __       \n| |  | |  / _ \\  | '_ \\      | __|       / __| |  _ \\   / _ \\  \\ \\ /\\ / /       \n| |__| | | (_) | | | | |     \\ |_        \\__ \\ | | | | | (_) |  \\ V  V /        \n|_____/   \\___/  |_| |_|      \\__|       |___/ |_| |_|  \\___/    \\_/\\_/         \n                                                                                \n                                                                                \n",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, cas := range cases {
		ret := methods.ProccessOutput(cas.input, graphics)
		if ret != cas.expected {
			fmt.Printf(Yellow+"Test => %s\n"+Reset, cas.name)
			t.Errorf(Red + "DONT PASS\n" + Reset)
		}
	}
}
