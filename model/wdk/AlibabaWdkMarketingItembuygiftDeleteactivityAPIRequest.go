package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest
删除买赠活动 API请求
alibaba.wdk.marketing.itembuygift.deleteactivity

删除买赠活动 */
type AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest struct {
	model.Params
	// 要删除的活动信息
	_param *CommonActivityParam
}

// New
