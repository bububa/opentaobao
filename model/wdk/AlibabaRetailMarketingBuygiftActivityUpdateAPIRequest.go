package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest
更新单品买赠活动 API请求
alibaba.retail.marketing.buygift.activity.update

同城零售单品买赠活动更新 */
type AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest struct {
	model.Params
	// 更新单品买赠活动参数
	_param *BuyGiftActivityOperateRequest
}

// New
