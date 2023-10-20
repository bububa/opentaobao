package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantidmixAPIResponse 商家用户id混淆 API返回值
// alibaba.tcls.aelophy.merchant.id.mix
//
// 商家用户id混淆
type AlibabatclsaelophymerchantidmixAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophymerchantidmixAPIResponseModel
}

// AlibabatclsaelophymerchantidmixAPIResponseModel is 商家用户id混淆 成功返回结果
type AlibabatclsaelophymerchantidmixAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_id_mix_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	ApiResult *AlibabatclsaelophymerchantidmixApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
