package security

import (
	"sync"
)

// Locale 结构体
type Locale struct {
	// 国家代码(参考ISO-3166)
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 语言代码(参考ISO-639)
	Language string `json:"language,omitempty" xml:"language,omitempty"`
}

var poolLocale = sync.Pool{
	New: func() any {
		return new(Locale)
	},
}

// GetLocale() 从对象池中获取Locale
func GetLocale() *Locale {
	return poolLocale.Get().(*Locale)
}

// ReleaseLocale 释放Locale
func ReleaseLocale(v *Locale) {
	v.Country = ""
	v.Language = ""
	poolLocale.Put(v)
}
