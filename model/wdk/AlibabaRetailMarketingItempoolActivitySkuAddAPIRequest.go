package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest
商品池活动新增商品 API请求
alibaba.retail.marketing.itempool.activity.sku.add

新增或更新商品池活动商品信息【同城零售】 */
type AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest struct {
	model.Params
	// 入参
	_param *ItemPoolActivityElementOperateRequest
}

// New
