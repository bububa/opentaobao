package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest
查询特价商品 API请求
alibaba.wdk.marketing.itemdiscount.queryitems

查询参加特价活动的商品优惠详情 */
type AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// New
