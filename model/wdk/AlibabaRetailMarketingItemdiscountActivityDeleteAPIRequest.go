package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest
删除单品特价活动【同城零售】 API请求
alibaba.retail.marketing.itemdiscount.activity.delete

同城零售单品特价活动删除 */
type AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest struct {
	model.Params
	// 删除活动参数
	_param *ItemDiscountActivityOperateRequest
}

// New
