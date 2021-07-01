package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest
查询买赠活动 API请求
alibaba.wdk.marketing.itembuygift.queryactivity

查询买赠活动 */
type AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest struct {
	model.Params
	// 查询入参
	_param *CommonActivityParam
}

// New
