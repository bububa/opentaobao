package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolActivityCreateAPIRequest
创建活动新接口 API请求
alibaba.wdk.marketing.itempool.activity.create

创建活动新接口，支持新工具玩法 */
type AlibabaWdkMarketingItempoolActivityCreateAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// New
