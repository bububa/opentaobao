package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolQueryactivityAPIRequest
查找商品池活动 API请求
alibaba.wdk.marketing.itempool.queryactivity

查找商品池活动 */
type AlibabaWdkMarketingItempoolQueryactivityAPIRequest struct {
	model.Params
	// 查询商品池活动入参
	_param0 *CommonActivityParam
}

// New
