package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNewretailPurchasePriceDeleteAPIRequest
共享库存 商户删除采购价 API请求
alibaba.newretail.purchase.price.delete

共享库存 商户删除采购价 */
type AlibabaNewretailPurchasePriceDeleteAPIRequest struct {
	model.Params
	// 调用入参
	_deletePurchasePriceRequest *DeletePurchasePriceRequest
}

// New
