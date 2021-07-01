package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemChangepriceQueryAPIRequest
按照价格变更时间段，查询会变更价格的单据的商品 API请求
alibaba.wdk.item.changeprice.query

*
     * 按照价格变更时间段，查询会变更价格的单据的商品
     * 传入QueryPriceChangeTypeEnum.BASE_PRICE,返回变价时间在startTime-endTime之间的所有单据
     * 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_START,
     * 返回活动开始时间在 startTime<=活动开始时间<endTime 之间的所有单品促销单据
     * 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_END,
     * 返回活动结束时间在 startTime<=活动结束时间<endTime 之间的所有单品促销单据
*/
type AlibabaWdkItemChangepriceQueryAPIRequest struct {
	model.Params
	// 变价的类型  * 查询变价的单据专用
	_type string
	// 开始时间
	_startTime string
	// 结束时间，结束时间-开始时间不能超过48小时
	_endTime string
	// 渠道店id
	_shopId string
}

// New
