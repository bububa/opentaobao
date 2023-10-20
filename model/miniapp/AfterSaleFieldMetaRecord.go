package miniapp

import (
	"sync"
)

// AfterSaleFieldMetaRecord 结构体
type AfterSaleFieldMetaRecord struct {
	// 子结构
	Children []string `json:"children,omitempty" xml:"children>string,omitempty"`
	// 字段名称
	FieldName string `json:"field_name,omitempty" xml:"field_name,omitempty"`
	// 字段描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 组件类型
	ComponentType string `json:"component_type,omitempty" xml:"component_type,omitempty"`
	// 配置信息
	CustomConfig string `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}

var poolAfterSaleFieldMetaRecord = sync.Pool{
	New: func() any {
		return new(AfterSaleFieldMetaRecord)
	},
}

// GetAfterSaleFieldMetaRecord() 从对象池中获取AfterSaleFieldMetaRecord
func GetAfterSaleFieldMetaRecord() *AfterSaleFieldMetaRecord {
	return poolAfterSaleFieldMetaRecord.Get().(*AfterSaleFieldMetaRecord)
}

// ReleaseAfterSaleFieldMetaRecord 释放AfterSaleFieldMetaRecord
func ReleaseAfterSaleFieldMetaRecord(v *AfterSaleFieldMetaRecord) {
	v.Children = v.Children[:0]
	v.FieldName = ""
	v.Description = ""
	v.ComponentType = ""
	v.CustomConfig = ""
	poolAfterSaleFieldMetaRecord.Put(v)
}
