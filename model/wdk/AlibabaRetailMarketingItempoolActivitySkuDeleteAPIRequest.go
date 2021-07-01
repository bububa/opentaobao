package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest
删除商品池活动商品【同城零售】 API请求
alibaba.retail.marketing.itempool.activity.sku.delete

删除商品池活动商品信息【同城零售】 */
type AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest struct {
	model.Params
	// 入参
	_param *ItemPoolActivityElementOperateRequest
}

// New
