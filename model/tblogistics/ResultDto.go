package tblogistics

// ResultDto 结构体
type ResultDto struct {
	// 返回核销订单列表
	WriteoffOrderList []WriteOffOrderDto `json:"writeoff_order_list,omitempty" xml:"writeoff_order_list>write_off_order_dto,omitempty"`
	// 发货提示文案
	ConsignDesc string `json:"consign_desc,omitempty" xml:"consign_desc,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
