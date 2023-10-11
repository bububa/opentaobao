package ascpchannel

// WaybillGenServ 结构体
type WaybillGenServ struct {
	// 服务类型：0为送装一体，1为只送到家不安装，2为只送不装到楼下，3为自提
	DeliveryType string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
}
