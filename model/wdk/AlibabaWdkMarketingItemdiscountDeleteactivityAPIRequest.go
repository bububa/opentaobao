package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest
删除商品特价活动 API请求
alibaba.wdk.marketing.itemdiscount.deleteactivity

删除商品特价活动 */
type AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityRequest
}

// New
