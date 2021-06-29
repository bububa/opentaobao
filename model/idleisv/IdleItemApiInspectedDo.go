package idleisv

// IdleItemApiInspectedDO 
type IdleItemApiInspectedDO struct {
    // 验货报告url链接(长度<=300)
    InspectReport   string `json:"inspect_report,omitempty" xml:"inspect_report,omitempty"`
    // 已验货入仓时间，时间戳，单位秒
    WareHouseTime   int64 `json:"ware_house_time,omitempty" xml:"ware_house_time,omitempty"`
    // 已验货入仓城市
    WareHouseCity   string `json:"ware_house_city,omitempty" xml:"ware_house_city,omitempty"`
}
