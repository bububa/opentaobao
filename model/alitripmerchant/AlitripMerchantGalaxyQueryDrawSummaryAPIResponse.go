package alitripmerchant

import (
	"encoding/xml"

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

// AlitripMerchantGalaxyQueryDrawSummaryAPIResponseModel is 星河-抽奖活动概要列表查询 成功返回结果
type AlitripMerchantGalaxyQueryDrawSummaryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_query_draw_summary_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyQueryDrawSummaryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
