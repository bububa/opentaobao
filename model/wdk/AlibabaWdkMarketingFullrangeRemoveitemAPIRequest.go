package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingFullrangeRemoveitemAPIRequest
全场活动删除购品 API请求
alibaba.wdk.marketing.fullrange.removeitem

删除换购商品 */
type AlibabaWdkMarketingFullrangeRemoveitemAPIRequest struct {
	model.Params
	// 商品sku信息
	_param0 *ItemStairSku
	// 活动信息
	_param1 *CommonActivityParam
}

// New
