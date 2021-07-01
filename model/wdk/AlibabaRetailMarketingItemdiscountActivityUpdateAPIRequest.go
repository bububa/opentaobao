package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest
更新单品特价活动【同城零售】 API请求
alibaba.retail.marketing.itemdiscount.activity.update

同城零售单品特价活动更新 */
type AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest struct {
	model.Params
	// 创建活动参数
	_param *ItemDiscountActivityOperateRequest
}

// New
