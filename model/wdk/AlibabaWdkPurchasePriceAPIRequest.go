package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkPurchasePriceAPIRequest
rt回传采购价 API请求
alibaba.wdk.purchase.price

猫超共享库存项目-rt回传采购价 */
type AlibabaWdkPurchasePriceAPIRequest struct {
	model.Params
	// 入参
	_wdkOpenPurchasePrice *WdkOpenPurchasePrice
}

// New
