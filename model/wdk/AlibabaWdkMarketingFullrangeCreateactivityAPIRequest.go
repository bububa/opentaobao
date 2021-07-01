package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingFullrangeCreateactivityAPIRequest
创建全场活动 API请求
alibaba.wdk.marketing.fullrange.createactivity

创建全场活动 */
type AlibabaWdkMarketingFullrangeCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *FullRangeActivity
}

// New
