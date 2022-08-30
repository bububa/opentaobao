package wlbimports

// HsfResult 结构体
type HsfResult struct {
	// data
	ExpressPreQueryResponseList []ExpressPreQueryResponse `json:"express_pre_query_response_list,omitempty" xml:"express_pre_query_response_list>express_pre_query_response,omitempty"`
	// 异常码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 网略异常码
	InternalErrorCode string `json:"internal_error_code,omitempty" xml:"internal_error_code,omitempty"`
	// 异常提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 响应业务对象
	AppointmentCancleReponse *AppointmentCancleReponse `json:"appointment_cancle_reponse,omitempty" xml:"appointment_cancle_reponse,omitempty"`
	// data
	AppointmentCreateResponse *AppointmentCreateResponse `json:"appointment_create_response,omitempty" xml:"appointment_create_response,omitempty"`
	// 响应体
	AppointmentOrderStatusResponse *AppointmentOrderStatusResponse `json:"appointment_order_status_response,omitempty" xml:"appointment_order_status_response,omitempty"`
	// 响应业务对象
	BigbagCancelResponse *BigbagCancelResponse `json:"bigbag_cancel_response,omitempty" xml:"bigbag_cancel_response,omitempty"`
	// data
	BigbagCreateResponse *BigbagCreateResponse `json:"bigbag_create_response,omitempty" xml:"bigbag_create_response,omitempty"`
	// 响应体
	Data *BigbagStatusResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 响应
	BigbagWaybillResponse *BigbagWaybillResponse `json:"bigbag_waybill_response,omitempty" xml:"bigbag_waybill_response,omitempty"`
	// 集货仓查询结果信息
	TransferStoreResponse *TransferStoreResponse `json:"transfer_store_response,omitempty" xml:"transfer_store_response,omitempty"`
	// 响应是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
