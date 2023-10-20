package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyQueryParticipateNumberAPIResponse 星河-抽奖活动次数查询 API返回值
// alitrip.merchant.galaxy.query.participate.number
//
// 雅高小程序抽奖活动，查询用户抽奖次数
type AlitripMerchantGalaxyQueryParticipateNumberAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyQueryParticipateNumberAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyQueryParticipateNumberAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyQueryParticipateNumberAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyQueryParticipateNumberAPIResponseModel is 星河-抽奖活动次数查询 成功返回结果
type AlitripMerchantGalaxyQueryParticipateNumberAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_query_participate_number_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyQueryParticipateNumberResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyQueryParticipateNumberAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyQueryParticipateNumberAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyQueryParticipateNumberAPIResponse)
	},
}

// GetAlitripMerchantGalaxyQueryParticipateNumberAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyQueryParticipateNumberAPIResponse
func GetAlitripMerchantGalaxyQueryParticipateNumberAPIResponse() *AlitripMerchantGalaxyQueryParticipateNumberAPIResponse {
	return poolAlitripMerchantGalaxyQueryParticipateNumberAPIResponse.Get().(*AlitripMerchantGalaxyQueryParticipateNumberAPIResponse)
}

// ReleaseAlitripMerchantGalaxyQueryParticipateNumberAPIResponse 将 AlitripMerchantGalaxyQueryParticipateNumberAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyQueryParticipateNumberAPIResponse(v *AlitripMerchantGalaxyQueryParticipateNumberAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyQueryParticipateNumberAPIResponse.Put(v)
}
