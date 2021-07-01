package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolDeleteactivityAPIRequest
删除商品池活动 API请求
alibaba.wdk.marketing.itempool.deleteactivity

删除商品池活动 */
type AlibabaWdkMarketingItempoolDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// New
