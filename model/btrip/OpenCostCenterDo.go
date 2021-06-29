package btrip

// OpenCostCenterDo 
type OpenCostCenterDo struct {
    // 成本中心编号
    Number   string `json:"number,omitempty" xml:"number,omitempty"`
    // 成本中心名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 商旅成本中心id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 商旅企业id
    Corpid   string `json:"corpid,omitempty" xml:"corpid,omitempty"`
}
