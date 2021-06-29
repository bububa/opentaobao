package inventory

// PromiseList 
type PromiseList struct {
    // 服务时效信息，发货时间比如2020-01-26；相对时间比如3，代表付款后3天内发货
    DeliveryTime   string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
    // 发货时间类型，0-绝对时间，1-相对时间
    DeliveryType   int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
    // 履约仓code
    PerformStore   string `json:"perform_store,omitempty" xml:"perform_store,omitempty"`
}
