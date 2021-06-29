package bus

// TopBusPriceAndStockUpdateRq 
type TopBusPriceAndStockUpdateRq struct {
    // 车次列表
    Numbers   []BusNumberInfoDTO `json:"numbers,omitempty" xml:"numbers>bus_number_info_dto,omitempty"`
}
