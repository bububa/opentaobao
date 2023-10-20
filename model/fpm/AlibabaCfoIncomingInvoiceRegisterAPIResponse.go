package fpm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfoIncomingInvoiceRegisterAPIResponse 发票登记接口 API返回值
// alibaba.cfo.incoming.invoice.register
//
// 发票登记接口
type AlibabaCfoIncomingInvoiceRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaCfoIncomingInvoiceRegisterAPIResponseModel
}

// AlibabaCfoIncomingInvoiceRegisterAPIResponseModel is 发票登记接口 成功返回结果
type AlibabaCfoIncomingInvoiceRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfo_incoming_invoice_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 响应消息
	ResponseMsg string `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
	// 发票Id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
