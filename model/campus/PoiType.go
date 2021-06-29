package campus

// PoiType 
type PoiType struct {
    // 类别名称
    TypeName   string `json:"type_name,omitempty" xml:"type_name,omitempty"`
    // 类别编码
    TypeCode   string `json:"type_code,omitempty" xml:"type_code,omitempty"`
    // 大类编码
    BigTypeCode   string `json:"big_type_code,omitempty" xml:"big_type_code,omitempty"`
    // 大类名称
    BigTypeName   string `json:"big_type_name,omitempty" xml:"big_type_name,omitempty"`
    // 大类ID
    BigTypeId   int64 `json:"big_type_id,omitempty" xml:"big_type_id,omitempty"`
    // 父类id
    Pid   int64 `json:"pid,omitempty" xml:"pid,omitempty"`
    // 类别的实际名称,空间或分组
    Classify   string `json:"classify,omitempty" xml:"classify,omitempty"`
    // 编码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // POI类型分类
    Category   int64 `json:"category,omitempty" xml:"category,omitempty"`
    // 是否删除
    IsDelete   bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
    // 描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 英文名称
    TypeEnName   string `json:"type_en_name,omitempty" xml:"type_en_name,omitempty"`
}
