package flightuppc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightinsuranceorderrefundAPIResponse 保险订单退保 API返回值
// alitrip.flight.insurance.order.refund
//
// 保险订单退保
type AlitripflightinsuranceorderrefundAPIResponse struct {
	model.CommonResponse
	AlitripflightinsuranceorderrefundAPIResponseModel
}

// AlitripflightinsuranceorderrefundAPIResponseModel is 保险订单退保 成功返回结果
type AlitripflightinsuranceorderrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_insurance_order_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息（调用成功时该字段为空）
	ErrMsgForClient string `json:"err_msg_for_client,omitempty" xml:"err_msg_for_client,omitempty"`
	// 保险订单号，，即tcOrderId
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
