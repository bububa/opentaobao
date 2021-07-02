package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNewretailPurchasePriceDeleteAPIRequest 共享库存 商户删除采购价 API请求
// alibaba.newretail.purchase.price.delete
//
// 共享库存 商户删除采购价
type AlibabaNewretailPurchasePriceDeleteAPIRequest struct {
	model.Params
	// 调用入参
	_deletePurchasePriceRequest *DeletePurchasePriceRequest
}

// NewAlibabaNewretailPurchasePriceDeleteRequest 初始化AlibabaNewretailPurchasePriceDeleteAPIRequest对象
func NewAlibabaNewretailPurchasePriceDeleteRequest() *AlibabaNewretailPurchasePriceDeleteAPIRequest {
	return &AlibabaNewretailPurchasePriceDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNewretailPurchasePriceDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.newretail.purchase.price.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNewretailPurchasePriceDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeletePurchasePriceRequest is DeletePurchasePriceRequest Setter
// 调用入参
func (r *AlibabaNewretailPurchasePriceDeleteAPIRequest) SetDeletePurchasePriceRequest(_deletePurchasePriceRequest *DeletePurchasePriceRequest) error {
	r._deletePurchasePriceRequest = _deletePurchasePriceRequest
	r.Set("delete_purchase_price_request", _deletePurchasePriceRequest)
	return nil
}

// GetDeletePurchasePriceRequest DeletePurchasePriceRequest Getter
func (r AlibabaNewretailPurchasePriceDeleteAPIRequest) GetDeletePurchasePriceRequest() *DeletePurchasePriceRequest {
	return r._deletePurchasePriceRequest
}
