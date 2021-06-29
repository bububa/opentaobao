package lstvending

// Range 
type Range struct {
    // 结束时间
    End   string `json:"end,omitempty" xml:"end,omitempty"`
    // 开始时间
    Begin   string `json:"begin,omitempty" xml:"begin,omitempty"`
}
