package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItempoolActivityCreateAPIRequest
创建商品池活动【同城零售】 API请求
alibaba.retail.marketing.itempool.activity.create

同城零售商品池活动创建 */
type AlibabaRetailMarketingItempoolActivityCreateAPIRequest struct {
	model.Params
	// 创建商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// New
