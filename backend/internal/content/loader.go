package content

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Loader struct {
	Root string
}

func NewLoader(root string) *Loader {
	return &Loader{Root: root}
}

func (l *Loader) Read(name string, target any) error {
	path := filepath.Join(l.Root, name)
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", name, err)
	}
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("unmarshal %s: %w", name, err)
	}
	return nil
}
