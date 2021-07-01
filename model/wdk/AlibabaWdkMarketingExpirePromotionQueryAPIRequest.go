package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingExpirePromotionQueryAPIRequest
短保优惠查询 API请求
alibaba.wdk.marketing.expire.promotion.query

短保优惠查询 */
type AlibabaWdkMarketingExpirePromotionQueryAPIRequest struct {
	model.Params
	// 店铺id
	_shopId string
	// 商品skucode
	_skuCode string
}

// New
