package retail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailelectroniccertificateconfirmAPIResponse 确认核销接口 API返回值
// alibaba.retail.electronic.certificate.confirm
//
// 确认核销接口
type AlibabaretailelectroniccertificateconfirmAPIResponse struct {
	model.CommonResponse
	AlibabaretailelectroniccertificateconfirmAPIResponseModel
}

// AlibabaretailelectroniccertificateconfirmAPIResponseModel is 确认核销接口 成功返回结果
type AlibabaretailelectroniccertificateconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_electronic_certificate_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaretailelectroniccertificateconfirmResult `json:"result,omitempty" xml:"result,omitempty"`
}
