package util

// LocaleDto 
type LocaleDto struct {
    // Locale编码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // Locale名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
