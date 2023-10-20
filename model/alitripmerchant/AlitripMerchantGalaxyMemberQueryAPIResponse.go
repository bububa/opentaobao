package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberQueryAPIResponse 星河-获取登录用户的信息 API返回值
// alitrip.merchant.galaxy.member.query
//
// 获取登录用户的信息
type AlitripMerchantGalaxyMemberQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMemberQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMemberQueryAPIResponseModel is 星河-获取登录用户的信息 成功返回结果
type AlitripMerchantGalaxyMemberQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMemberQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMemberQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMemberQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMemberQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMemberQueryAPIResponse
func GetAlitripMerchantGalaxyMemberQueryAPIResponse() *AlitripMerchantGalaxyMemberQueryAPIResponse {
	return poolAlitripMerchantGalaxyMemberQueryAPIResponse.Get().(*AlitripMerchantGalaxyMemberQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMemberQueryAPIResponse 将 AlitripMerchantGalaxyMemberQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberQueryAPIResponse(v *AlitripMerchantGalaxyMemberQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberQueryAPIResponse.Put(v)
}
