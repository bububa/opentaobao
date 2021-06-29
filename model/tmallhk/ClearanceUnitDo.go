package tmallhk

// ClearanceUnitDo 
type ClearanceUnitDo struct {
    // 第一数量，最多4位小数
    FirstQuantity   string `json:"first_quantity,omitempty" xml:"first_quantity,omitempty"`
    // 第一单位，单位编码
    FirstUnit   string `json:"first_unit,omitempty" xml:"first_unit,omitempty"`
    // 第二数量，最多4位小数
    SecondQuantity   string `json:"second_quantity,omitempty" xml:"second_quantity,omitempty"`
    // 第二单位，单位编码
    SecondUnit   string `json:"second_unit,omitempty" xml:"second_unit,omitempty"`
}
