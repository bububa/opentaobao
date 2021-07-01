package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest
删除单品买赠活动 API请求
alibaba.retail.marketing.buygift.activity.delete

同城零售单品特价活动删除 */
type AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest struct {
	model.Params
	// 删除活动参数
	_param *ItemDiscountActivityOperateRequest
}

// New
