package cli

import (
	"os"
	"path/filepath"
)


type CLI struct{
	rootDir string
	outputFileName string
	outputFile *os.File
	customProhibitedFilesDir string
	prohibitedFiles map[string]interface{}
}

func New() *CLI{
	root, _ := filepath.Abs(".")
	return &CLI{
		rootDir: root,
		outputFileName: "result.md",
		customProhibitedFilesDir: "",
	}
}

func (c *CLI) Open() error{
	outputFile, err := os.Create(c.outputFileName)
	if err != nil {
		return err 
	}
	c.outputFile = outputFile

	c.convertProhFilesToMap()
	return nil
}

func (c *CLI) Close(){
	c.outputFile.Close()
}

// Functions use before open
func (c *CLI) SetDir(dir string) error{
	root, err := filepath.Abs(dir)
	if err != nil{
		return err
	}

	c.rootDir = root
	return nil
}
func (c *CLI) SetOutputFileName(name string){
	name += ".md"
	c.outputFileName = name
}
func (c *CLI) SetCustomProhibit(name string){
	c.customProhibitedFilesDir = name
}