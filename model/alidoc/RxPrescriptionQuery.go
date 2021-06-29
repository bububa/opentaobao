package alidoc

// RxPrescriptionQuery 
type RxPrescriptionQuery struct {
    // 订单号多个以逗号分开
    BizOrderIds   string `json:"biz_order_ids,omitempty" xml:"biz_order_ids,omitempty"`
}
