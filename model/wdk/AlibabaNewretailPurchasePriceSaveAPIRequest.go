package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNewretailPurchasePriceSaveAPIRequest
共享库存 采购价上传接口 API请求
alibaba.newretail.purchase.price.save

共享库存业务 供应商上传商品采购价 */
type AlibabaNewretailPurchasePriceSaveAPIRequest struct {
	model.Params
	// 接口入参
	_savePurchasePriceRequest *SavePurchasePriceRequest
}

// New
