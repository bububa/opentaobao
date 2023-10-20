package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkpurchasepriceAPIRequest rt回传采购价 API请求
// alibaba.wdk.purchase.price
//
// 猫超共享库存项目-rt回传采购价
type AlibabawdkpurchasepriceAPIRequest struct {
	model.Params
	// 入参
	_wdkOpenPurchasePrice *WdkOpenPurchasePrice
}

// NewAlibabawdkpurchasepriceRequest 初始化AlibabawdkpurchasepriceAPIRequest对象
func NewAlibabawdkpurchasepriceRequest() *AlibabawdkpurchasepriceAPIRequest {
	return &AlibabawdkpurchasepriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkpurchasepriceAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.purchase.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkpurchasepriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkpurchasepriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWdkOpenPurchasePrice is WdkOpenPurchasePrice Setter
// 入参
func (r *AlibabawdkpurchasepriceAPIRequest) SetWdkOpenPurchasePrice(_wdkOpenPurchasePrice *WdkOpenPurchasePrice) error {
	r._wdkOpenPurchasePrice = _wdkOpenPurchasePrice
	r.Set("wdk_open_purchase_price", _wdkOpenPurchasePrice)
	return nil
}

// GetWdkOpenPurchasePrice WdkOpenPurchasePrice Getter
func (r AlibabawdkpurchasepriceAPIRequest) GetWdkOpenPurchasePrice() *WdkOpenPurchasePrice {
	return r._wdkOpenPurchasePrice
}
