package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest
创建商品特价活动 API请求
alibaba.wdk.marketing.itemdiscount.createactivity

创建商品特价活动 */
type AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemDiscountActivityRequest
}

// New
