package tmallhk

// CtsCarriage 
type CtsCarriage struct {
    // 托运开始时间
    Begin   string `json:"begin,omitempty" xml:"begin,omitempty"`
    // 托运单号
    CarriageNo   string `json:"carriage_no,omitempty" xml:"carriage_no,omitempty"`
    // 托运公司名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 托运结束时间
    End   string `json:"end,omitempty" xml:"end,omitempty"`
}
