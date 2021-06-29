package wdk

// Consignee 
type Consignee struct {
    // 收件人名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 收件人电话
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 配送地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 配送坐标
    Geo   string `json:"geo,omitempty" xml:"geo,omitempty"`
    // 配送开始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 配送结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 0:高德坐标系，1:其他坐标系（需要坐标修正）
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 城市名称，仅在type不为0的情况下传入
    City   string `json:"city,omitempty" xml:"city,omitempty"`
}
