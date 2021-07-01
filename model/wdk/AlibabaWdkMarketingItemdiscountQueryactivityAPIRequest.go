package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest
查找特价活动 API请求
alibaba.wdk.marketing.itemdiscount.queryactivity

查找特价活动 */
type AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *CommonActivityParam
}

// New
