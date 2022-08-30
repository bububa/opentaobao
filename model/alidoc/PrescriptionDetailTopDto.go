package alidoc

// PrescriptionDetailTopDto 结构体
type PrescriptionDetailTopDto struct {
	// 订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 处方图地址
	PrescriptionPicUrl string `json:"prescription_pic_url,omitempty" xml:"prescription_pic_url,omitempty"`
}
