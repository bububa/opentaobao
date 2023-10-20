package flightuppc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightinsuranceordermodifyAPIResponse 保险订单批改申请 API返回值
// alitrip.flight.insurance.order.modify
//
// 保险订单批改申请
type AlitripflightinsuranceordermodifyAPIResponse struct {
	model.CommonResponse
	AlitripflightinsuranceordermodifyAPIResponseModel
}

// AlitripflightinsuranceordermodifyAPIResponseModel is 保险订单批改申请 成功返回结果
type AlitripflightinsuranceordermodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_insurance_order_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息（调用成功时该字段为空）
	ErrMsgForClient string `json:"err_msg_for_client,omitempty" xml:"err_msg_for_client,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
