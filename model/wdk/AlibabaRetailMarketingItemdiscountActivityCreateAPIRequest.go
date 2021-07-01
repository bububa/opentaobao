package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest
创建单品特价活动【同城零售】 API请求
alibaba.retail.marketing.itemdiscount.activity.create

同城零售单品特价活动创建 */
type AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest struct {
	model.Params
	// 创建活动参数
	_param *ItemDiscountActivityOperateRequest
}

// New
