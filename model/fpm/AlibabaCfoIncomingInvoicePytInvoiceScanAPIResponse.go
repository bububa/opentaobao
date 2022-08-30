package fpm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfoIncomingInvoicePytInvoiceScanAPIResponse 票易通发票ocr信息同步 API返回值
// alibaba.cfo.incoming.invoice.pyt.invoice.scan
//
// 票易通发票ocr信息同步
type AlibabaCfoIncomingInvoicePytInvoiceScanAPIResponse struct {
	model.CommonResponse
	AlibabaCfoIncomingInvoicePytInvoiceScanAPIResponseModel
}

// AlibabaCfoIncomingInvoicePytInvoiceScanAPIResponseModel is 票易通发票ocr信息同步 成功返回结果
type AlibabaCfoIncomingInvoicePytInvoiceScanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfo_incoming_invoice_pyt_invoice_scan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 空值
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 响应码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 响应消息
	ResponseMsg string `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
	// 是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
