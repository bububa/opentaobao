package bus

// CityChangeRq 
type CityChangeRq struct {
    // 代理商ID
    AgentId   int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
    // 城市列表属性
    List   []BusCityChangeDTO `json:"list,omitempty" xml:"list>bus_city_change_dto,omitempty"`
}
