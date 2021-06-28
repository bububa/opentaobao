package bus

// CityChangeRq 
/* model for simplify = false
type CityChangeRq struct {

    // 代理商ID
    
    AgentId   int64 `json:"agent_id,omitempty"`
    

    // 城市列表属性
    
    List  struct {
        BusCityChangeDto  []BusCityChangeDto `json:"bus_city_change_dto,omitempty"`
    } `json:"list,omitempty"`
    

}
*/

// CityChangeRq 
type CityChangeRq struct {

    // 代理商ID
    AgentId   int64 `json:"agent_id,omitempty"`

    // 城市列表属性
    List   []BusCityChangeDto `json:"list,omitempty"`

}
