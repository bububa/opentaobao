package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest
删除单品买赠活动商品 API请求
alibaba.retail.marketing.buygift.activity.sku.delete

删除单品买赠活动商品信息【同城零售】 */
type AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest struct {
	model.Params
	// 删除买赠活动商品参数
	_param *BuyGiftActivitySkuOperateRequest
}

// New
