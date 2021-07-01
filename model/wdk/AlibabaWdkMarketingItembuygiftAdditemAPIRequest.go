package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItembuygiftAdditemAPIRequest
增加买赠活动商品。【注意，此接口暂不支持并发！】 API请求
alibaba.wdk.marketing.itembuygift.additem

增加买赠活动商品。【注意，此接口暂不支持并发！】 */
type AlibabaWdkMarketingItembuygiftAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemBuyGiftSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// New
