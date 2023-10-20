package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemChangepriceQuery 按照价格变更时间段，查询会变更价格的单据的商品
// alibaba.wdk.item.changeprice.query
//
// *
//   - 按照价格变更时间段，查询会变更价格的单据的商品
//   - 传入QueryPriceChangeTypeEnum.BASE_PRICE,返回变价时间在startTime-endTime之间的所有单据
//   - 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_START,
//   - 返回活动开始时间在 startTime&lt;=活动开始时间&lt;endTime 之间的所有单品促销单据
//   - 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_END,
//   - 返回活动结束时间在 startTime&lt;=活动结束时间&lt;endTime 之间的所有单品促销单据
func AlibabaWdkItemChangepriceQuery(clt *core.SDKClient, req *wdk.AlibabaWdkItemChangepriceQueryAPIRequest, resp *wdk.AlibabaWdkItemChangepriceQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
