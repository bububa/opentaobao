package campus

import (
	"sync"
)

// TypeAttrInstance 结构体
type TypeAttrInstance struct {
	// attrName
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// attrCode
	AttrCode string `json:"attr_code,omitempty" xml:"attr_code,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// valueType
	ValueType string `json:"value_type,omitempty" xml:"value_type,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// typeAttrRefId
	TypeAttrRefId int64 `json:"type_attr_ref_id,omitempty" xml:"type_attr_ref_id,omitempty"`
}

var poolTypeAttrInstance = sync.Pool{
	New: func() any {
		return new(TypeAttrInstance)
	},
}

// GetTypeAttrInstance() 从对象池中获取TypeAttrInstance
func GetTypeAttrInstance() *TypeAttrInstance {
	return poolTypeAttrInstance.Get().(*TypeAttrInstance)
}

// ReleaseTypeAttrInstance 释放TypeAttrInstance
func ReleaseTypeAttrInstance(v *TypeAttrInstance) {
	v.AttrName = ""
	v.AttrCode = ""
	v.Value = ""
	v.ValueType = ""
	v.Uuid = ""
	v.TypeAttrRefId = 0
	poolTypeAttrInstance.Put(v)
}
