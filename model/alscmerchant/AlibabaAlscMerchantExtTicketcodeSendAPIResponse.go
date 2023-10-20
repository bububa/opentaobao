package alscmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscmerchantextticketcodesendAPIResponse 异步发码 API返回值
// alibaba.alsc.merchant.ext.ticketcode.send
//
// 外部券异步发码
type AlibabaalscmerchantextticketcodesendAPIResponse struct {
	model.CommonResponse
	AlibabaalscmerchantextticketcodesendAPIResponseModel
}

// AlibabaalscmerchantextticketcodesendAPIResponseModel is 异步发码 成功返回结果
type AlibabaalscmerchantextticketcodesendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_merchant_ext_ticketcode_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求id
	TicketRequestId string `json:"ticket_request_id,omitempty" xml:"ticket_request_id,omitempty"`
	// 该字段用于描述本次返回中的业务属性，现有：BIZ_ALREADY_SUCCESS（幂等业务码）。
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}
