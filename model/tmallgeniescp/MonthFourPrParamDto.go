package tmallgeniescp

// MonthFourPrParamDTO 
type MonthFourPrParamDTO struct {
    // 关键值日期
    KeyFigureDate   string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
    // PR数量
    PrQty   string `json:"pr_qty,omitempty" xml:"pr_qty,omitempty"`
    // 产品线
    ProductionLine   string `json:"production_line,omitempty" xml:"production_line,omitempty"`
    // 扩展参数
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    // 租户
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}
