package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingBuygiftActivityCreateAPIRequest
创建买赠活动 API请求
alibaba.retail.marketing.buygift.activity.create

同城供给买赠活动创建 */
type AlibabaRetailMarketingBuygiftActivityCreateAPIRequest struct {
	model.Params
	// 创建活动参数
	_param *BuyGiftActivityOperateRequest
}

// New
