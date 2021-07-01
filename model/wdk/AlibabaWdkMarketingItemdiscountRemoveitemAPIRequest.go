package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest
移除报名的商品 API请求
alibaba.wdk.marketing.itemdiscount.removeitem

将报名特价活动的商品从特价活动中移除 */
type AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemDiscountSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// New
