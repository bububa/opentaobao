package bus

// TopBusNumberUpdateRq 
/* model for simplify = false
type TopBusNumberUpdateRq struct {

    // 车次列表
    
    Numbers  struct {
        BusNumberDto  []BusNumberDto `json:"bus_number_dto,omitempty"`
    } `json:"numbers,omitempty"`
    

}
*/

// TopBusNumberUpdateRq 
type TopBusNumberUpdateRq struct {

    // 车次列表
    Numbers   []BusNumberDto `json:"numbers,omitempty"`

}
