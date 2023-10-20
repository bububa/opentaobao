package campus

import (
	"sync"
)

// PropertyApiDto 结构体
type PropertyApiDto struct {
	// 参数点名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 参数点编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 参数点类型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 值类型名称
	ValueTypeName string `json:"value_type_name,omitempty" xml:"value_type_name,omitempty"`
	// 元数据编码
	PropertyKindCode string `json:"property_kind_code,omitempty" xml:"property_kind_code,omitempty"`
	// 控制类参数枚举值
	ControlEnumValue string `json:"control_enum_value,omitempty" xml:"control_enum_value,omitempty"`
	// 单位编码
	UnitCode string `json:"unit_code,omitempty" xml:"unit_code,omitempty"`
	// 参数点id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 参数点类型id
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 值类型id
	ValueType int64 `json:"value_type,omitempty" xml:"value_type,omitempty"`
	// 元数据id
	PropertyKind int64 `json:"property_kind,omitempty" xml:"property_kind,omitempty"`
	// 单位id
	UnitId int64 `json:"unit_id,omitempty" xml:"unit_id,omitempty"`
	// 是否报警类参数
	Alarm bool `json:"alarm,omitempty" xml:"alarm,omitempty"`
}

var poolPropertyApiDto = sync.Pool{
	New: func() any {
		return new(PropertyApiDto)
	},
}

// GetPropertyApiDto() 从对象池中获取PropertyApiDto
func GetPropertyApiDto() *PropertyApiDto {
	return poolPropertyApiDto.Get().(*PropertyApiDto)
}

// ReleasePropertyApiDto 释放PropertyApiDto
func ReleasePropertyApiDto(v *PropertyApiDto) {
	v.Name = ""
	v.Code = ""
	v.TypeName = ""
	v.ValueTypeName = ""
	v.PropertyKindCode = ""
	v.ControlEnumValue = ""
	v.UnitCode = ""
	v.Id = 0
	v.Type = 0
	v.ValueType = 0
	v.PropertyKind = 0
	v.UnitId = 0
	v.Alarm = false
	poolPropertyApiDto.Put(v)
}
