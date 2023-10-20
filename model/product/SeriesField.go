package product

import (
	"sync"
)

// SeriesField 结构体
type SeriesField struct {
	// 文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// type
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolSeriesField = sync.Pool{
	New: func() any {
		return new(SeriesField)
	},
}

// GetSeriesField() 从对象池中获取SeriesField
func GetSeriesField() *SeriesField {
	return poolSeriesField.Get().(*SeriesField)
}

// ReleaseSeriesField 释放SeriesField
func ReleaseSeriesField(v *SeriesField) {
	v.Text = ""
	v.Value = ""
	v.Key = ""
	v.Type = ""
	poolSeriesField.Put(v)
}
