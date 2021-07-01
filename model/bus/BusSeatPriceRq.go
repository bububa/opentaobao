package bus

// BusSeatPriceRq 结构体
type BusSeatPriceRq struct {
	// 车次ID
	ScheduleId string `json:"schedule_id,omitempty" xml:"schedule_id,omitempty"`
}
