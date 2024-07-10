package cli


import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (c *CLI) RecursiveCopyFiles() error {
	err := filepath.Walk(c.rootDir, func(path string, info os.FileInfo, err error) error {
		// Mac feature - hidden ds_store file
		base := filepath.Base(path)
		if info.IsDir() && (strings.HasPrefix(base, ".") || base == ".DS_Store") {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			if c.ifProhibitedExtension(base) || c.isProhibitedFile(base) {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			c.outputFile.WriteString(fmt.Sprintf("\n### %s\n", base))
			c.outputFile.WriteString("\n```\n")
			io.Copy(c.outputFile, file)
			c.outputFile.WriteString("\n```\n")
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *CLI) StructureToFile() error{
	c.outputFile.WriteString("\n## Структура проекта:\n")
	c.outputFile.WriteString("\n```\n Если у файла нет расширения - это каталог ")
	c.outputFile.WriteString("под каталогом если у файла/папки есть -, это указывает на его вложенность вышележащий файл ")
	c.outputFile.WriteString("количество -(тире) указывает на уровень вложенность \n```\n")
	var structure strings.Builder
	c.RecursiveCopyStructure(c.rootDir, 0, &structure)
	c.outputFile.WriteString("\n```\n")
	c.outputFile.WriteString(structure.String())
	c.outputFile.WriteString("\n```\n")
	return nil
}
func (c *CLI) RecursiveCopyStructure(dir string, depth int, st *strings.Builder) error{
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		indent := strings.Repeat("-", depth)
		name := entry.Name()

		if c.ifProhibitedExtension(name) || c.isProhibitedFile(name){
			continue
		}

		st.WriteString(fmt.Sprintf("%s%s\n", indent, name))

		if entry.IsDir() {
			subdir := filepath.Join(dir, name)
			c.RecursiveCopyStructure(subdir, depth+1, st)
		}
	}

	return nil
}