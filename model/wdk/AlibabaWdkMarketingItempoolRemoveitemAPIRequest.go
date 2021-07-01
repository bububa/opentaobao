package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolRemoveitemAPIRequest
移除商品池里面的商品 API请求
alibaba.wdk.marketing.itempool.removeitem

移除商品池里面的商品 */
type AlibabaWdkMarketingItempoolRemoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// New
