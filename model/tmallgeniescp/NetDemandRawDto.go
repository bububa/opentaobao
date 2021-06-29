package tmallgeniescp

// NetDemandRawDto 
type NetDemandRawDto struct {

    // 扩展参数
    
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    

    // 租户
    
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    

    // 关键日期值
    
    KeyFigureDate   string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
    

    // 地点
    
    Locid   string `json:"locid,omitempty" xml:"locid,omitempty"`
    

    // 二级物料净需求
    
    NetDemandRaw   int64 `json:"net_demand_raw,omitempty" xml:"net_demand_raw,omitempty"`
    

    // 二级物料-物料编码
    
    MaterialCode   string `json:"material_code,omitempty" xml:"material_code,omitempty"`
    

}
