package bus

// TopBusPriceAndStockUpdateRq 结构体
type TopBusPriceAndStockUpdateRq struct {
	// 车次列表
	Numbers []BusNumberInfoDto `json:"numbers,omitempty" xml:"numbers>bus_number_info_dto,omitempty"`
}
