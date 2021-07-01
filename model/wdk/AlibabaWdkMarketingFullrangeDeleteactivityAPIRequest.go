package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest
全场活动删除活动接口 API请求
alibaba.wdk.marketing.fullrange.deleteactivity

全场活动删除活动 */
type AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// New
