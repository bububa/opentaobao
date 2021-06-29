package nlife

// Deliverlist 
type Deliverlist struct {
    // 发货单的批次号
    BatchNo   int64 `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
    // 发货单号
    ConsignNo   string `json:"consign_no,omitempty" xml:"consign_no,omitempty"`
    // 物流公司
    LogisticsCompany   string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
    // 物流单号
    LogisticsNo   string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
    // 发货时间
    GmtConsign   string `json:"gmt_consign,omitempty" xml:"gmt_consign,omitempty"`
}
