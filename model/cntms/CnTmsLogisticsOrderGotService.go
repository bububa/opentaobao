package cntms

// CnTmsLogisticsOrderGotService 
type CnTmsLogisticsOrderGotService struct {
    // 揽收日期yyyyMMdd
    GotDate   string `json:"got_date,omitempty" xml:"got_date,omitempty"`
    // 揽收时间段 09:00-10:00
    GotRange   string `json:"got_range,omitempty" xml:"got_range,omitempty"`
}
