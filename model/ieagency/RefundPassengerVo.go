package ieagency

// RefundPassengerVo 
type RefundPassengerVo struct {
    // 乘机人ID
    PassenerId   int64 `json:"passener_id,omitempty" xml:"passener_id,omitempty"`
    // 乘机人姓名
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    // 乘机人类型(Adult(0, "成人"),     Child(1, "儿童"),     StudentAbroad(2, "留学生"),     Infant(3, "婴儿")
    PassengerType   int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
}
