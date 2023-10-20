package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyVerifySignatureAPIResponse 微信账号生物认证 API返回值
// alitrip.merchant.galaxy.verify.signature
//
// 在用户注册等场景下，如果账号风险等级过高，需要进行生物认证
type AlitripMerchantGalaxyVerifySignatureAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyVerifySignatureAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyVerifySignatureAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyVerifySignatureAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyVerifySignatureAPIResponseModel is 微信账号生物认证 成功返回结果
type AlitripMerchantGalaxyVerifySignatureAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_verify_signature_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyVerifySignatureResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyVerifySignatureAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyVerifySignatureAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyVerifySignatureAPIResponse)
	},
}

// GetAlitripMerchantGalaxyVerifySignatureAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyVerifySignatureAPIResponse
func GetAlitripMerchantGalaxyVerifySignatureAPIResponse() *AlitripMerchantGalaxyVerifySignatureAPIResponse {
	return poolAlitripMerchantGalaxyVerifySignatureAPIResponse.Get().(*AlitripMerchantGalaxyVerifySignatureAPIResponse)
}

// ReleaseAlitripMerchantGalaxyVerifySignatureAPIResponse 将 AlitripMerchantGalaxyVerifySignatureAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyVerifySignatureAPIResponse(v *AlitripMerchantGalaxyVerifySignatureAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyVerifySignatureAPIResponse.Put(v)
}
