package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest
添加单品买赠活动商品 API请求
alibaba.retail.marketing.buygift.activity.sku.add

新增或更新单品买赠活动商品信息【同城零售】 */
type AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *BuyGiftActivitySkuOperateRequest
}

// New
