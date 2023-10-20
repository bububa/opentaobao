package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyQueryDrawSummaryAPIResponse 星河-抽奖活动概要列表查询 API返回值
// alitrip.merchant.galaxy.query.draw.summary
//
// 雅高小程序抽奖活动列表概要查询
type AlitripMerchantGalaxyQueryDrawSummaryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyQueryDrawSummaryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyQueryDrawSummaryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyQueryDrawSummaryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyQueryDrawSummaryAPIResponseModel is 星河-抽奖活动概要列表查询 成功返回结果
type AlitripMerchantGalaxyQueryDrawSummaryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_query_draw_summary_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyQueryDrawSummaryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyQueryDrawSummaryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyQueryDrawSummaryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyQueryDrawSummaryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyQueryDrawSummaryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyQueryDrawSummaryAPIResponse
func GetAlitripMerchantGalaxyQueryDrawSummaryAPIResponse() *AlitripMerchantGalaxyQueryDrawSummaryAPIResponse {
	return poolAlitripMerchantGalaxyQueryDrawSummaryAPIResponse.Get().(*AlitripMerchantGalaxyQueryDrawSummaryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyQueryDrawSummaryAPIResponse 将 AlitripMerchantGalaxyQueryDrawSummaryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyQueryDrawSummaryAPIResponse(v *AlitripMerchantGalaxyQueryDrawSummaryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyQueryDrawSummaryAPIResponse.Put(v)
}
