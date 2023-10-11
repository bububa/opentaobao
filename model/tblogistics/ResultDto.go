package tblogistics

// ResultDto 结构体
type ResultDto struct {
	// 返回核销订单列表
	WriteoffOrderList []WriteOffOrderDto `json:"writeoff_order_list,omitempty" xml:"writeoff_order_list>write_off_order_dto,omitempty"`
	// -
	Consign *ConsignDto `json:"consign,omitempty" xml:"consign,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
