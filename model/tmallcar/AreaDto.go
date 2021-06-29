package tmallcar

// AreaDTO 
type AreaDTO struct {
    // 公司详细地址
    CompanyAdress   string `json:"company_adress,omitempty" xml:"company_adress,omitempty"`
    // 公司email
    CompanyEmail   string `json:"company_email,omitempty" xml:"company_email,omitempty"`
    // 公司名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 电话
    CompanyPhone   string `json:"company_phone,omitempty" xml:"company_phone,omitempty"`
    // 行政区域地址,菜鸟地址库2级或者3级地址
    DivisionId   int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
    // 行政区名称
    DivisionName   string `json:"division_name,omitempty" xml:"division_name,omitempty"`
    // 状态:1正常,2删除
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
