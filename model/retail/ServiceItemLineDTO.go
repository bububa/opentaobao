package retail

// ServiceItemLineDto 结构体
type ServiceItemLineDto struct {
	// 服务商品编码
	ServiceSpuCode string `json:"service_spu_code,omitempty" xml:"service_spu_code,omitempty"`
	// 服务子订单id
	ServiceId int64 `json:"service_id,omitempty" xml:"service_id,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
