package sqlhandler

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type SQLHandler struct {
	db          *gorm.DB
	templateDir string
}

func New(db *gorm.DB, templateDir string) *SQLHandler {
	return &SQLHandler{
		db:          db,
		templateDir: templateDir,
	}
}

func (h *SQLHandler) LoadSQLTemplate(fileName string) (string, error) {
	filePath := filepath.Join(h.templateDir, fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read SQL template: %w", err)
	}

	content = regexp.MustCompile(`\s+`).ReplaceAll(content, []byte(" "))
	content = regexp.MustCompile(`\s*,\s*`).ReplaceAll(content, []byte(","))

	return string(content), nil
}

func (h *SQLHandler) ReplaceNamedVariables(sqlTemplate string, params map[string]interface{}) (string, []interface{}) {
	placeholders := []interface{}{}

	re := regexp.MustCompile(`\@\w+`)
	matches := re.FindAllString(sqlTemplate, -1)

	for _, match := range matches {
		paramName := strings.TrimPrefix(match, "@")
		if value, exists := params[paramName]; exists {
			sqlTemplate = strings.ReplaceAll(sqlTemplate, match, "?")
			placeholders = append(placeholders, value)
		}
	}

	return sqlTemplate, placeholders
}

func (h *SQLHandler) ExecuteSQLWithNamedParams(fileName string, params map[string]interface{}, result interface{}) error {
	sqlTemplate, err := h.LoadSQLTemplate(fileName)
	if err != nil {
		return err
	}

	query, args := h.ReplaceNamedVariables(sqlTemplate, params)

	if result == nil {
		return h.db.Exec(query, args...).Error
	}

	return h.db.Raw(query, args...).Scan(result).Error
}
