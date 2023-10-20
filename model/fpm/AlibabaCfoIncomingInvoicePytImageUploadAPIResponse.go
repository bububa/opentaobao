package fpm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfoincominginvoicepytimageuploadAPIResponse 票易通发票影像上传 API返回值
// alibaba.cfo.incoming.invoice.pyt.image.upload
//
// 票易通发票影像上传
type AlibabacfoincominginvoicepytimageuploadAPIResponse struct {
	model.CommonResponse
	AlibabacfoincominginvoicepytimageuploadAPIResponseModel
}

// AlibabacfoincominginvoicepytimageuploadAPIResponseModel is 票易通发票影像上传 成功返回结果
type AlibabacfoincominginvoicepytimageuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfo_incoming_invoice_pyt_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 文件resourceId
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 响应码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 相应消息
	ResponseMsg string `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
	// 是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
