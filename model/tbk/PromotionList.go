package tbk

// PromotionList 结构体
type PromotionList struct {
	// 权益开始时间，精确到毫秒时间戳
	Entryusedstarttime string `json:"entry_used_start_time,omitempty" xml:"entry_used_start_time,omitempty"`
	// 权益结束时间，精确到毫秒时间戳
	Entryusedendtime string `json:"entry_used_end_time,omitempty" xml:"entry_used_end_time,omitempty"`
	// 权益起用门槛，满X元可用，券场景为满元，精确到分，如满100元可用
	Entrycondition string `json:"entry_condition,omitempty" xml:"entry_condition,omitempty"`
	// 权益面额，券场景为减钱，精确到分
	Entrydiscount string `json:"entry_discount,omitempty" xml:"entry_discount,omitempty"`
}
