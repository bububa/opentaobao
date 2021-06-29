package alihealth2

// Ingredients 
type Ingredients struct {
    // 成分单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // 成分数值
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 成分名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
