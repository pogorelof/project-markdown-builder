package cli

import (
	"bufio"
	"os"
	"strings"
)

var (
	defaultProhibitedFilesDir = "./configs/prohibited.txt"
)

func (c *CLI) convertProhFilesToMap() error {
	c.prohibitedFiles = map[string]interface{}{
		"ext":   make(map[string]interface{}),
		"files": make(map[string]interface{}),
	}

	filesMap := c.prohibitedFiles["files"].(map[string]interface{})
	filesMap[c.outputFileName] = ""

	err := c.convert(defaultProhibitedFilesDir)
	if err != nil {
		return err
	}
	if c.customProhibitedFilesDir != "" {
		err := c.convert(c.customProhibitedFilesDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CLI) convert(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#") {
			extMap := c.prohibitedFiles["ext"].(map[string]interface{})
			extMap[text[1:]] = ""
			continue
		}

		filesMap := c.prohibitedFiles["files"].(map[string]interface{})
		filesMap[text] = ""
	}
	return nil
}

func (c *CLI) isProhibitedFile(name string) bool {
	filesMap := c.prohibitedFiles["files"].(map[string]interface{})
	if _, ok := filesMap[name]; ok {
		return true
	}
	return false
}

func (c *CLI) ifProhibitedExtension(name string) bool {
	extMap := c.prohibitedFiles["ext"].(map[string]interface{})
	for k := range extMap {
		if strings.HasSuffix(name, k) {
			return true
		}
	}
	return false
}
