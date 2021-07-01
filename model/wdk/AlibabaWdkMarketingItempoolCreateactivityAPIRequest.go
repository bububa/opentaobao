package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolCreateactivityAPIRequest
添加商品池活动 API请求
alibaba.wdk.marketing.itempool.createactivity

添加商品池活动 */
type AlibabaWdkMarketingItempoolCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// New
