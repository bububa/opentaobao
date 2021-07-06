package bus

// CityChangeRq 结构体
type CityChangeRq struct {
	// 城市列表属性
	List []BusCityChangeDto `json:"list,omitempty" xml:"list>bus_city_change_dto,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}
