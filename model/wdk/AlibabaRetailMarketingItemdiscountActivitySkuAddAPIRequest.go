package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest
特价活动新增商品 API请求
alibaba.retail.marketing.itemdiscount.activity.sku.add

新增或更新活动商品信息【同城零售】 */
type AlibabaRetailMarketingItemdiscountActivitySkuAddAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *ItemDiscountActivityElementOperateRequest
}

// New
