package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNewretailPurchasePriceSaveAPIRequest 共享库存 采购价上传接口 API请求
// alibaba.newretail.purchase.price.save
//
// 共享库存业务 供应商上传商品采购价
type AlibabaNewretailPurchasePriceSaveAPIRequest struct {
	model.Params
	// 接口入参
	_savePurchasePriceRequest *SavePurchasePriceRequest
}

// NewAlibabaNewretailPurchasePriceSaveRequest 初始化AlibabaNewretailPurchasePriceSaveAPIRequest对象
func NewAlibabaNewretailPurchasePriceSaveRequest() *AlibabaNewretailPurchasePriceSaveAPIRequest {
	return &AlibabaNewretailPurchasePriceSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNewretailPurchasePriceSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.newretail.purchase.price.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNewretailPurchasePriceSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SavePurchasePriceRequest Setter
// 接口入参
func (r *AlibabaNewretailPurchasePriceSaveAPIRequest) SetSavePurchasePriceRequest(_savePurchasePriceRequest *SavePurchasePriceRequest) error {
	r._savePurchasePriceRequest = _savePurchasePriceRequest
	r.Set("save_purchase_price_request", _savePurchasePriceRequest)
	return nil
}

// Get SavePurchasePriceRequest Getter
func (r AlibabaNewretailPurchasePriceSaveAPIRequest) GetSavePurchasePriceRequest() *SavePurchasePriceRequest {
	return r._savePurchasePriceRequest
}
