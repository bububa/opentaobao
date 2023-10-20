package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNewretailPurchasePriceSaveAPIRequest) Reset() {
	r._savePurchasePriceRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNewretailPurchasePriceSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.newretail.purchase.price.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNewretailPurchasePriceSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNewretailPurchasePriceSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSavePurchasePriceRequest is SavePurchasePriceRequest Setter
// 接口入参
func (r *AlibabaNewretailPurchasePriceSaveAPIRequest) SetSavePurchasePriceRequest(_savePurchasePriceRequest *SavePurchasePriceRequest) error {
	r._savePurchasePriceRequest = _savePurchasePriceRequest
	r.Set("save_purchase_price_request", _savePurchasePriceRequest)
	return nil
}

// GetSavePurchasePriceRequest SavePurchasePriceRequest Getter
func (r AlibabaNewretailPurchasePriceSaveAPIRequest) GetSavePurchasePriceRequest() *SavePurchasePriceRequest {
	return r._savePurchasePriceRequest
}

var poolAlibabaNewretailPurchasePriceSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNewretailPurchasePriceSaveRequest()
	},
}

// GetAlibabaNewretailPurchasePriceSaveRequest 从 sync.Pool 获取 AlibabaNewretailPurchasePriceSaveAPIRequest
func GetAlibabaNewretailPurchasePriceSaveAPIRequest() *AlibabaNewretailPurchasePriceSaveAPIRequest {
	return poolAlibabaNewretailPurchasePriceSaveAPIRequest.Get().(*AlibabaNewretailPurchasePriceSaveAPIRequest)
}

// ReleaseAlibabaNewretailPurchasePriceSaveAPIRequest 将 AlibabaNewretailPurchasePriceSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaNewretailPurchasePriceSaveAPIRequest(v *AlibabaNewretailPurchasePriceSaveAPIRequest) {
	v.Reset()
	poolAlibabaNewretailPurchasePriceSaveAPIRequest.Put(v)
}
