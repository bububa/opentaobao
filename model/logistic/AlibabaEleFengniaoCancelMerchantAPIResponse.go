package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoCancelMerchantAPIResponse 商户取消 API返回值
// alibaba.ele.fengniao.cancel.merchant
//
// 商户取消配送
type AlibabaEleFengniaoCancelMerchantAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoCancelMerchantAPIResponseModel
}

// AlibabaEleFengniaoCancelMerchantAPIResponseModel is 商户取消 成功返回结果
type AlibabaEleFengniaoCancelMerchantAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_cancel_merchant_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
