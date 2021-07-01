package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest
删除特价活动商品 API请求
alibaba.retail.marketing.itemdiscount.activity.sku.delete

删除活动商品信息【同城零售】 */
type AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *ItemDiscountActivityElementOperateRequest
}

// New
