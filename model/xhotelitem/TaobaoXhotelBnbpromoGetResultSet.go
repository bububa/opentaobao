package xhotelitem

// TaobaoXhotelBnbpromoGetResultSet 结构体
type TaobaoXhotelBnbpromoGetResultSet struct {
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 民宿活动信息
	BnbPromo *BnbPromoDto `json:"bnb_promo,omitempty" xml:"bnb_promo,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
