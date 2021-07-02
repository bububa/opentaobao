package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPurchasePriceAPIRequest rt回传采购价 API请求
// alibaba.wdk.purchase.price
//
// 猫超共享库存项目-rt回传采购价
type AlibabaWdkPurchasePriceAPIRequest struct {
	model.Params
	// 入参
	_wdkOpenPurchasePrice *WdkOpenPurchasePrice
}

// NewAlibabaWdkPurchasePriceRequest 初始化AlibabaWdkPurchasePriceAPIRequest对象
func NewAlibabaWdkPurchasePriceRequest() *AlibabaWdkPurchasePriceAPIRequest {
	return &AlibabaWdkPurchasePriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPurchasePriceAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.purchase.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPurchasePriceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWdkOpenPurchasePrice is WdkOpenPurchasePrice Setter
// 入参
func (r *AlibabaWdkPurchasePriceAPIRequest) SetWdkOpenPurchasePrice(_wdkOpenPurchasePrice *WdkOpenPurchasePrice) error {
	r._wdkOpenPurchasePrice = _wdkOpenPurchasePrice
	r.Set("wdk_open_purchase_price", _wdkOpenPurchasePrice)
	return nil
}

// GetWdkOpenPurchasePrice WdkOpenPurchasePrice Getter
func (r AlibabaWdkPurchasePriceAPIRequest) GetWdkOpenPurchasePrice() *WdkOpenPurchasePrice {
	return r._wdkOpenPurchasePrice
}
