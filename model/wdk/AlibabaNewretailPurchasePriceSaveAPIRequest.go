package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanewretailpurchasepricesaveAPIRequest 共享库存 采购价上传接口 API请求
// alibaba.newretail.purchase.price.save
//
// 共享库存业务 供应商上传商品采购价
type AlibabanewretailpurchasepricesaveAPIRequest struct {
	model.Params
	// 接口入参
	_savePurchasePriceRequest *SavePurchasePriceRequest
}

// NewAlibabanewretailpurchasepricesaveRequest 初始化AlibabanewretailpurchasepricesaveAPIRequest对象
func NewAlibabanewretailpurchasepricesaveRequest() *AlibabanewretailpurchasepricesaveAPIRequest {
	return &AlibabanewretailpurchasepricesaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanewretailpurchasepricesaveAPIRequest) GetApiMethodName() string {
	return "alibaba.newretail.purchase.price.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanewretailpurchasepricesaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanewretailpurchasepricesaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSavePurchasePriceRequest is SavePurchasePriceRequest Setter
// 接口入参
func (r *AlibabanewretailpurchasepricesaveAPIRequest) SetSavePurchasePriceRequest(_savePurchasePriceRequest *SavePurchasePriceRequest) error {
	r._savePurchasePriceRequest = _savePurchasePriceRequest
	r.Set("save_purchase_price_request", _savePurchasePriceRequest)
	return nil
}

// GetSavePurchasePriceRequest SavePurchasePriceRequest Getter
func (r AlibabanewretailpurchasepricesaveAPIRequest) GetSavePurchasePriceRequest() *SavePurchasePriceRequest {
	return r._savePurchasePriceRequest
}
