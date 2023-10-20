package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanewretailpurchasepricedeleteAPIRequest 共享库存 商户删除采购价 API请求
// alibaba.newretail.purchase.price.delete
//
// 共享库存 商户删除采购价
type AlibabanewretailpurchasepricedeleteAPIRequest struct {
	model.Params
	// 调用入参
	_deletePurchasePriceRequest *DeletePurchasePriceRequest
}

// NewAlibabanewretailpurchasepricedeleteRequest 初始化AlibabanewretailpurchasepricedeleteAPIRequest对象
func NewAlibabanewretailpurchasepricedeleteRequest() *AlibabanewretailpurchasepricedeleteAPIRequest {
	return &AlibabanewretailpurchasepricedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanewretailpurchasepricedeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.newretail.purchase.price.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanewretailpurchasepricedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanewretailpurchasepricedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeletePurchasePriceRequest is DeletePurchasePriceRequest Setter
// 调用入参
func (r *AlibabanewretailpurchasepricedeleteAPIRequest) SetDeletePurchasePriceRequest(_deletePurchasePriceRequest *DeletePurchasePriceRequest) error {
	r._deletePurchasePriceRequest = _deletePurchasePriceRequest
	r.Set("delete_purchase_price_request", _deletePurchasePriceRequest)
	return nil
}

// GetDeletePurchasePriceRequest DeletePurchasePriceRequest Getter
func (r AlibabanewretailpurchasepricedeleteAPIRequest) GetDeletePurchasePriceRequest() *DeletePurchasePriceRequest {
	return r._deletePurchasePriceRequest
}
