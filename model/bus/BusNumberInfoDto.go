package bus

// BusNumberInfoDto 
type BusNumberInfoDto struct {

    // 代理商id
    AgentId   int64 `json:"agent_id,omitempty"`

    // 车次id
    ScheduleId   string `json:"schedule_id,omitempty"`

    // 出发时间
    DepartTime   string `json:"depart_time,omitempty"`

    // 库存
    Stock   int64 `json:"stock,omitempty"`

    // 价格
    FullPrice   int64 `json:"full_price,omitempty"`

    // 出发城市名称
    CityName   string `json:"city_name,omitempty"`

}
