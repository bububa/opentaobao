package tmallgeniescp

// SupplierForecastRawParamDto 
type SupplierForecastRawParamDto struct {

    // 扩展参数
    
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    

    // 租户
    
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    

    // 关键日期值
    
    KeyFigureDate   string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
    

    // 计划量
    
    Forecast   string `json:"forecast,omitempty" xml:"forecast,omitempty"`
    

    // 地点 - 二级物料供应商code
    
    LocationCode   string `json:"location_code,omitempty" xml:"location_code,omitempty"`
    

    // 二级物料-物料编码
    
    PrdId   string `json:"prd_id,omitempty" xml:"prd_id,omitempty"`
    

    // 地点 - 成品供应商仓库&二级物料中央仓
    
    LocationCodeTo   string `json:"location_code_to,omitempty" xml:"location_code_to,omitempty"`
    

}
