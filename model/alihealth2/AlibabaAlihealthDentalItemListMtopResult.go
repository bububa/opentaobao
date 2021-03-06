package alihealth2

// AlibabaAlihealthDentalItemListMtopResult 结构体
type AlibabaAlihealthDentalItemListMtopResult struct {
	// model
	Goods []NormalGoodsVo `json:"goods,omitempty" xml:"goods>normal_goods_vo,omitempty"`
	// 200
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
