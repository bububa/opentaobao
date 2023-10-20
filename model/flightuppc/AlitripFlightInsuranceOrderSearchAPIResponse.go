package flightuppc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightinsuranceordersearchAPIResponse 查询保险订单详情 API返回值
// alitrip.flight.insurance.order.search
//
// 查询保险订单详情
type AlitripflightinsuranceordersearchAPIResponse struct {
	model.CommonResponse
	AlitripflightinsuranceordersearchAPIResponseModel
}

// AlitripflightinsuranceordersearchAPIResponseModel is 查询保险订单详情 成功返回结果
type AlitripflightinsuranceordersearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_insurance_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 保险订单
	InsOrders []InsOrderOpenDto `json:"ins_orders,omitempty" xml:"ins_orders>ins_order_open_dto,omitempty"`
	// 错误信息，调用成功是该字段为空
	ErrMsgForClient string `json:"err_msg_for_client,omitempty" xml:"err_msg_for_client,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
