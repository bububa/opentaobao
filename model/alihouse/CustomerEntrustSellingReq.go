package alihouse

// CustomerEntrustSellingReq 结构体
type CustomerEntrustSellingReq struct {
	// 委托状态
	EntrustedStatus int64 `json:"entrusted_status,omitempty" xml:"entrusted_status,omitempty"`
	// 管家ID
	ButlerServiceId int64 `json:"butler_service_id,omitempty" xml:"butler_service_id,omitempty"`
	// 委托业务主ID
	EntrustSellingId int64 `json:"entrust_selling_id,omitempty" xml:"entrust_selling_id,omitempty"`
}
