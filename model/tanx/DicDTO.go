package tanx

// DicDTO 
type DicDTO struct {
    // 数据项ID
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 数据项值
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
}
