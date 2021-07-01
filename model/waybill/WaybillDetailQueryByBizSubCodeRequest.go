package waybill

// WaybillDetailQueryByBizSubCodeRequest 结构体
type WaybillDetailQueryByBizSubCodeRequest struct {
	// 订单号
	BizSubCode string `json:"biz_sub_code,omitempty" xml:"biz_sub_code,omitempty"`
	// 请求id
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
}
