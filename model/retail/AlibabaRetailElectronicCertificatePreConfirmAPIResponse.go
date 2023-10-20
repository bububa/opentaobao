package retail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailelectroniccertificatepreconfirmAPIResponse 贩卖机开始核销接口 API返回值
// alibaba.retail.electronic.certificate.pre.confirm
//
// 零售终端贩卖机开始核销接口,返回待领的商品ID
type AlibabaretailelectroniccertificatepreconfirmAPIResponse struct {
	model.CommonResponse
	AlibabaretailelectroniccertificatepreconfirmAPIResponseModel
}

// AlibabaretailelectroniccertificatepreconfirmAPIResponseModel is 贩卖机开始核销接口 成功返回结果
type AlibabaretailelectroniccertificatepreconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_electronic_certificate_pre_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaretailelectroniccertificatepreconfirmResult `json:"result,omitempty" xml:"result,omitempty"`
}
