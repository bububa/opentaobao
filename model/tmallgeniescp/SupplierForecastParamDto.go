package tmallgeniescp

// SupplierForecastParamDto 
type SupplierForecastParamDto struct {
    // 扩展参数
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    // 租户
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    // 关键日期
    KeyFigureDate   string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
    // 供应商预测数量
    Forecast   string `json:"forecast,omitempty" xml:"forecast,omitempty"`
    // 地点
    LocId   string `json:"loc_id,omitempty" xml:"loc_id,omitempty"`
    // 物料号
    PrdId   string `json:"prd_id,omitempty" xml:"prd_id,omitempty"`
}
