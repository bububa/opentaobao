package tmallgeniescp

// SupplierFeedbackDTO 
type SupplierFeedbackDTO struct {
    // 扩展参数
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    // 租户
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    // 关键值日期
    KeyFigureDate   string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
    // 承诺反馈数量
    CommitQty   string `json:"commit_qty,omitempty" xml:"commit_qty,omitempty"`
    // 地点
    LocId   string `json:"loc_id,omitempty" xml:"loc_id,omitempty"`
    // 物料号
    PrdId   string `json:"prd_id,omitempty" xml:"prd_id,omitempty"`
    // 地点 - 二级物料供应商
    LocationCode   string `json:"location_code,omitempty" xml:"location_code,omitempty"`
    // 地点 - 采购分仓目的loc
    LocationCodeTo   string `json:"location_code_to,omitempty" xml:"location_code_to,omitempty"`
}
