package travel

// StructItinerary 
type StructItinerary struct {

    // 必填，行程序号，标识是第几天的行程
    DayOrder   int64 `json:"day_order,omitempty"`

    // 必填，当天行程包含的多个活动信息
    Activities   []ItineraryActivity `json:"activities,omitempty"`

}
