package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItempoolActivityUpdateAPIRequest
更新商品池活动【同城零售】 API请求
alibaba.retail.marketing.itempool.activity.update

同城零售商品池活动更新 */
type AlibabaRetailMarketingItempoolActivityUpdateAPIRequest struct {
	model.Params
	// 更新商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// New
