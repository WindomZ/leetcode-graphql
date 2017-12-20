package leetcodegraphql

import "strings"

// Code the struct of leetcode codes.
type Code struct {
	Text        string `json:"text"`
	Value       string `json:"value"`
	DefaultCode string `json:"defaultCode"`
}

// Codes the slice of Code
type Codes []*Code

// Code returns Code with lang.
func (cs Codes) Code(lang string) *Code {
	for _, c := range cs {
		if strings.ToLower(c.Text) == lang || c.Value == lang {
			return c
		}
	}
	return nil
}
