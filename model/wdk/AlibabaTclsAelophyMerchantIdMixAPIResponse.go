package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantIdMixAPIResponse 商家用户id混淆 API返回值
// alibaba.tcls.aelophy.merchant.id.mix
//
// 商家用户id混淆
type AlibabaTclsAelophyMerchantIdMixAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantIdMixAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantIdMixAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantIdMixAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantIdMixAPIResponseModel is 商家用户id混淆 成功返回结果
type AlibabaTclsAelophyMerchantIdMixAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_id_mix_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	ApiResult *AlibabaTclsAelophyMerchantIdMixApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantIdMixAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantIdMixAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantIdMixAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantIdMixAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantIdMixAPIResponse
func GetAlibabaTclsAelophyMerchantIdMixAPIResponse() *AlibabaTclsAelophyMerchantIdMixAPIResponse {
	return poolAlibabaTclsAelophyMerchantIdMixAPIResponse.Get().(*AlibabaTclsAelophyMerchantIdMixAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantIdMixAPIResponse 将 AlibabaTclsAelophyMerchantIdMixAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantIdMixAPIResponse(v *AlibabaTclsAelophyMerchantIdMixAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantIdMixAPIResponse.Put(v)
}
