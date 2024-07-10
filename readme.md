# Project Markdown Builder
Project Markdown Builder - is a tool that collects all your project files into one markdown file. The tool collects the file in a convenient form to provide a AI for project analysis. The output file has the structure of the project and a copy of all the files below. 

## Quick Start
1. `git clone https://github.com/pogorelof/project-markdown-builder.git`
2. `cd project-markdown-builder `
3. `make build` (or `go build ./cmd/cli`)
4. Copy the compiled file to your project directory
5. Run it! Result will be in result.md 

## Flags
- `-dir` - select project directory. "." by  default
- `-o` - change name of output file(just name, without extension). "result" by  default 
- `-proh` - directory of file containing names of prohibited files.

## Prohibited files
The tool contains default banned files in the directory `configs/prohibited.txt`. You can add a file that complements it. Just create a txt file and enter from a new line the files that should be ignored. If you need to ignore all files with a certain extension, enter # and then the write this extension like `#.md`

