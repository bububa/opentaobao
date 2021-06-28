package bus

// TopBusPriceAndStockUpdateRq 
/* model for simplify = false
type TopBusPriceAndStockUpdateRq struct {

    // 车次列表
    
    Numbers  struct {
        BusNumberInfoDto  []BusNumberInfoDto `json:"bus_number_info_dto,omitempty"`
    } `json:"numbers,omitempty"`
    

}
*/

// TopBusPriceAndStockUpdateRq 
type TopBusPriceAndStockUpdateRq struct {

    // 车次列表
    Numbers   []BusNumberInfoDto `json:"numbers,omitempty"`

}
