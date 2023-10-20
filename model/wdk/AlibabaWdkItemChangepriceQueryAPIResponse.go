package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemchangepricequeryAPIResponse 按照价格变更时间段，查询会变更价格的单据的商品 API返回值
// alibaba.wdk.item.changeprice.query
//
// *
//   - 按照价格变更时间段，查询会变更价格的单据的商品
//   - 传入QueryPriceChangeTypeEnum.BASE_PRICE,返回变价时间在startTime-endTime之间的所有单据
//   - 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_START,
//   - 返回活动开始时间在 startTime&lt;=活动开始时间&lt;endTime 之间的所有单品促销单据
//   - 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_END,
//   - 返回活动结束时间在 startTime&lt;=活动结束时间&lt;endTime 之间的所有单品促销单据
type AlibabawdkitemchangepricequeryAPIResponse struct {
	model.CommonResponse
	AlibabawdkitemchangepricequeryAPIResponseModel
}

// AlibabawdkitemchangepricequeryAPIResponseModel is 按照价格变更时间段，查询会变更价格的单据的商品 成功返回结果
type AlibabawdkitemchangepricequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_changeprice_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabawdkitemchangepricequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
