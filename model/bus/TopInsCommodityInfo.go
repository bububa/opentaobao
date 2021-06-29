package bus

// TopInsCommodityInfo 
type TopInsCommodityInfo struct {
    // 票价，单位：分
    TicketPrice   int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
    // 出行日期，精确到分钟 yyyyMMddHHmm
    TravelDate   string `json:"travel_date,omitempty" xml:"travel_date,omitempty"`
    // 起始站点
    StartStationId   string `json:"start_station_id,omitempty" xml:"start_station_id,omitempty"`
    // 行程时长，精确到分钟
    ItineraryTime   int64 `json:"itinerary_time,omitempty" xml:"itinerary_time,omitempty"`
}
