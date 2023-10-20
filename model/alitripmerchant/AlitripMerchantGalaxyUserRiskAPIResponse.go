package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyUserRiskAPIResponse 查询微信账号的风险等级 API返回值
// alitrip.merchant.galaxy.user.risk
//
// 【星河服务】查询微信账号的风险等级
type AlitripMerchantGalaxyUserRiskAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyUserRiskAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyUserRiskAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyUserRiskAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyUserRiskAPIResponseModel is 查询微信账号的风险等级 成功返回结果
type AlitripMerchantGalaxyUserRiskAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_user_risk_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyUserRiskResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyUserRiskAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyUserRiskAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyUserRiskAPIResponse)
	},
}

// GetAlitripMerchantGalaxyUserRiskAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyUserRiskAPIResponse
func GetAlitripMerchantGalaxyUserRiskAPIResponse() *AlitripMerchantGalaxyUserRiskAPIResponse {
	return poolAlitripMerchantGalaxyUserRiskAPIResponse.Get().(*AlitripMerchantGalaxyUserRiskAPIResponse)
}

// ReleaseAlitripMerchantGalaxyUserRiskAPIResponse 将 AlitripMerchantGalaxyUserRiskAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyUserRiskAPIResponse(v *AlitripMerchantGalaxyUserRiskAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyUserRiskAPIResponse.Put(v)
}
