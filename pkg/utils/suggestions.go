package utils

import "github.com/c-bata/go-prompt"

// suggestions list
var Suggestions = []prompt.Suggest{
	// General
	{"exit", "Exit adretctl"},
	{"help", "get help method"},

	//debug mode
	{"+x", "see which command is launched, equivalent of bash +x"},
	{"-x", "disable debug mode"},

	//linux utility
	{"pwd", "get current directory"},
	{"rm", "remove a file or directory (-f available for directory)"},
	{"grep", "print lines that match patterns (-R and -i available)"},
	{"cat", "concatenate files and print on the standard output"},
	{"ls", "list directory contents ( -a, -l, -R available"},
	{"tree", "list contents of directories in a tree-like format"},
	{"cp", "copy files and directories"},
	{"hostname", "show the system's host name"},
	{"cd", "change working directory (accept ~ and - arguments)"},
	{"env", "get environnement variable"},
}
