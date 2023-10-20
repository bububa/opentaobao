package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopaimainftcertificateapplycallbackAPIResponse 数字藏品版权证书申请结果回调 API返回值
// taobao.paimai.nft.certificate.applycallback
//
// 数字藏品版权证书申请结果回调
type TaobaopaimainftcertificateapplycallbackAPIResponse struct {
	model.CommonResponse
	TaobaopaimainftcertificateapplycallbackAPIResponseModel
}

// TaobaopaimainftcertificateapplycallbackAPIResponseModel is 数字藏品版权证书申请结果回调 成功返回结果
type TaobaopaimainftcertificateapplycallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"paimai_nft_certificate_applycallback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID 方便排查问题记录
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 错误信息
	FailMessage string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
	// true/false 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
