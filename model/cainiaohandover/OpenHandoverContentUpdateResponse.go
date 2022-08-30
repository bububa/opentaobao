package cainiaohandover

// OpenHandoverContentUpdateResponse 结构体
type OpenHandoverContentUpdateResponse struct {
	// 报错小包列表
	UpdateErrorParcelOrderList []HandoverContentUpdateErrorParcelDto `json:"update_error_parcel_order_list,omitempty" xml:"update_error_parcel_order_list>handover_content_update_error_parcel_dto,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
