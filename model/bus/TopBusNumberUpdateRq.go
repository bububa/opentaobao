package bus

// TopBusNumberUpdateRq 结构体
type TopBusNumberUpdateRq struct {
	// 车次列表
	Numbers []BusNumberDto `json:"numbers,omitempty" xml:"numbers>bus_number_dto,omitempty"`
}
