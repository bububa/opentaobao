package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest
创建买赠活动 API请求
alibaba.wdk.marketing.itembuygift.createactivity

创建买赠活动 */
type AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemBuyGiftActivity
}

// New
