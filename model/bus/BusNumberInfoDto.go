package bus

// BusNumberInfoDto 结构体
type BusNumberInfoDto struct {
	// 代理商id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 车次id
	ScheduleId string `json:"schedule_id,omitempty" xml:"schedule_id,omitempty"`
	// 出发时间
	DepartTime string `json:"depart_time,omitempty" xml:"depart_time,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 价格
	FullPrice int64 `json:"full_price,omitempty" xml:"full_price,omitempty"`
	// 出发城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}
