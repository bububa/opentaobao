package campus

// TypeAttrInstance 结构体
type TypeAttrInstance struct {
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// valueType
	ValueType string `json:"value_type,omitempty" xml:"value_type,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// typeAttrRefId
	TypeAttrRefId int64 `json:"type_attr_ref_id,omitempty" xml:"type_attr_ref_id,omitempty"`
	// attrName
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// attrCode
	AttrCode string `json:"attr_code,omitempty" xml:"attr_code,omitempty"`
}
