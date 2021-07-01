package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest
查询买赠活动下的商品 API请求
alibaba.wdk.marketing.itembuygift.queryitems

查询买赠活动下的商品 */
type AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// New
