package tmallgeniescp

// C2MConsensusDemandParamDTO 
type C2MConsensusDemandParamDTO struct {
    // 关键日期值
    KeyFigureDate   string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
    // c2m需求数量
    C2mConsensusDemand   string `json:"c2m_consensus_demand,omitempty" xml:"c2m_consensus_demand,omitempty"`
    // 渠道ID
    CustId   string `json:"cust_id,omitempty" xml:"cust_id,omitempty"`
    // 物料ID
    PrdId   string `json:"prd_id,omitempty" xml:"prd_id,omitempty"`
    // 扩展参数
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    // 租户
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    // c2m分配数量
    C2mQuantity   string `json:"c2m_quantity,omitempty" xml:"c2m_quantity,omitempty"`
}
