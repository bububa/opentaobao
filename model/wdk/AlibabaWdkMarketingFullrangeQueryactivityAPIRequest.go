package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingFullrangeQueryactivityAPIRequest
全场活动查询活动 API请求
alibaba.wdk.marketing.fullrange.queryactivity

全场活动查询活动 */
type AlibabaWdkMarketingFullrangeQueryactivityAPIRequest struct {
	model.Params
	// 查询活动入参
	_param0 *CommonActivityParam
}

// New
