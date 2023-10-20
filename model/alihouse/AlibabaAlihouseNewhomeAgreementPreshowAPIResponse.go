package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeagreementpreshowAPIResponse 预览地址获取接口 API返回值
// alibaba.alihouse.newhome.agreement.preshow
//
// 预览地址获取接口
type AlibabaalihousenewhomeagreementpreshowAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeagreementpreshowAPIResponseModel
}

// AlibabaalihousenewhomeagreementpreshowAPIResponseModel is 预览地址获取接口 成功返回结果
type AlibabaalihousenewhomeagreementpreshowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_agreement_preshow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaalihousenewhomeagreementpreshowResult `json:"result,omitempty" xml:"result,omitempty"`
}
