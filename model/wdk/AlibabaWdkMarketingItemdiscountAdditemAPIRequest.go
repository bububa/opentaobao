package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItemdiscountAdditemAPIRequest
报名特价商品 API请求
alibaba.wdk.marketing.itemdiscount.additem

在商品特价活动中报名特价商品 */
type AlibabaWdkMarketingItemdiscountAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemDiscountSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// New
