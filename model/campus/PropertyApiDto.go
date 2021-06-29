package campus

// PropertyApiDTO 
type PropertyApiDTO struct {
    // 参数点id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 参数点名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 参数点编码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 参数点类型id
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 参数点类型名称
    TypeName   string `json:"type_name,omitempty" xml:"type_name,omitempty"`
    // 是否报警类参数
    Alarm   bool `json:"alarm,omitempty" xml:"alarm,omitempty"`
    // 值类型id
    ValueType   int64 `json:"value_type,omitempty" xml:"value_type,omitempty"`
    // 值类型名称
    ValueTypeName   string `json:"value_type_name,omitempty" xml:"value_type_name,omitempty"`
    // 元数据id
    PropertyKind   int64 `json:"property_kind,omitempty" xml:"property_kind,omitempty"`
    // 元数据编码
    PropertyKindCode   string `json:"property_kind_code,omitempty" xml:"property_kind_code,omitempty"`
    // 控制类参数枚举值
    ControlEnumValue   string `json:"control_enum_value,omitempty" xml:"control_enum_value,omitempty"`
    // 单位id
    UnitId   int64 `json:"unit_id,omitempty" xml:"unit_id,omitempty"`
    // 单位编码
    UnitCode   string `json:"unit_code,omitempty" xml:"unit_code,omitempty"`
}
