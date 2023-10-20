package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundCsapplyrenderAPIResponse 商家代客售后逆向申请渲染获取 API返回值
// alibaba.tcls.aelophy.refund.csapplyrender
//
// 提供商家代客售后逆向申请渲染获取的接口
type AlibabaTclsAelophyRefundCsapplyrenderAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyRefundCsapplyrenderAPIResponseModel
}

// AlibabaTclsAelophyRefundCsapplyrenderAPIResponseModel is 商家代客售后逆向申请渲染获取 成功返回结果
type AlibabaTclsAelophyRefundCsapplyrenderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_csapplyrender_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	ApiResult *AlibabaTclsAelophyRefundCsapplyrenderApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
