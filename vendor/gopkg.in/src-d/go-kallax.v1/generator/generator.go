// Package generator implements the processor of source code and generator of
// kallax models based on Go source code.
package generator // import "gopkg.in/src-d/go-kallax.v1/generator"

import (
	"fmt"
	"os"
	"runtime/debug"
)

// Generator is in charge of generating files for packages.
type Generator struct {
	filename string
}

// NewGenerator creates a new generator that can save on the given filename.
func NewGenerator(filename string) *Generator {
	return &Generator{filename}
}

// Generate writes the file with the contents of the given package.
func (g *Generator) Generate(pkg *Package) error {
	return g.writeFile(pkg)
}

func (g *Generator) writeFile(pkg *Package) (err error) {
	file, err := os.Create(g.filename)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("kallax: PANIC during '%s' generation:\n%s\n\n", g.filename, r)
			if err == nil {
				err = fmt.Errorf(string(debug.Stack()))
			}
		}

		file.Close()
		if err != nil {
			if os.Remove(g.filename) == nil {
				fmt.Println("kallax: No file generated due to an occurred error:")
			} else {
				fmt.Printf("kallax: The autogenerated file '%s' could not be completed nor deleted due to an occurred error:\n", g.filename)
			}
		}
	}()

	return Base.Execute(file, pkg)
}