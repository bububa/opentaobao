package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDropshippingStoreSaveAPIRequest 阿里巴巴dropshipping店铺数据保存接口 API请求
// alibaba.dropshipping.store.save
//
// 阿里巴巴dropshipping店铺数据保存
type AlibabaDropshippingStoreSaveAPIRequest struct {
	model.Params
	// store type
	_storeType string
	// store url
	_storeUrl string
}

// NewAlibabaDropshippingStoreSaveRequest 初始化AlibabaDropshippingStoreSaveAPIRequest对象
func NewAlibabaDropshippingStoreSaveRequest() *AlibabaDropshippingStoreSaveAPIRequest {
	return &AlibabaDropshippingStoreSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingStoreSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.dropshipping.store.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingStoreSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDropshippingStoreSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreType is StoreType Setter
// store type
func (r *AlibabaDropshippingStoreSaveAPIRequest) SetStoreType(_storeType string) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// GetStoreType StoreType Getter
func (r AlibabaDropshippingStoreSaveAPIRequest) GetStoreType() string {
	return r._storeType
}

// SetStoreUrl is StoreUrl Setter
// store url
func (r *AlibabaDropshippingStoreSaveAPIRequest) SetStoreUrl(_storeUrl string) error {
	r._storeUrl = _storeUrl
	r.Set("store_url", _storeUrl)
	return nil
}

// GetStoreUrl StoreUrl Getter
func (r AlibabaDropshippingStoreSaveAPIRequest) GetStoreUrl() string {
	return r._storeUrl
}
