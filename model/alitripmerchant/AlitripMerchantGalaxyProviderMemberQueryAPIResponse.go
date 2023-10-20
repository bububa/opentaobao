package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyProviderMemberQueryAPIResponse 提供会员查询接口 API返回值
// alitrip.merchant.galaxy.provider.member.query
//
// 星河产品=提供会查询服务
type AlitripMerchantGalaxyProviderMemberQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyProviderMemberQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyProviderMemberQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyProviderMemberQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyProviderMemberQueryAPIResponseModel is 提供会员查询接口 成功返回结果
type AlitripMerchantGalaxyProviderMemberQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_provider_member_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyProviderMemberQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyProviderMemberQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyProviderMemberQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyProviderMemberQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyProviderMemberQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyProviderMemberQueryAPIResponse
func GetAlitripMerchantGalaxyProviderMemberQueryAPIResponse() *AlitripMerchantGalaxyProviderMemberQueryAPIResponse {
	return poolAlitripMerchantGalaxyProviderMemberQueryAPIResponse.Get().(*AlitripMerchantGalaxyProviderMemberQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyProviderMemberQueryAPIResponse 将 AlitripMerchantGalaxyProviderMemberQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyProviderMemberQueryAPIResponse(v *AlitripMerchantGalaxyProviderMemberQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyProviderMemberQueryAPIResponse.Put(v)
}
