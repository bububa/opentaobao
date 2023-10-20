package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadropshippingstoresaveAPIRequest 阿里巴巴dropshipping店铺数据保存接口 API请求
// alibaba.dropshipping.store.save
//
// 阿里巴巴dropshipping店铺数据保存
type AlibabadropshippingstoresaveAPIRequest struct {
	model.Params
	// store type
	_storeType string
	// store url
	_storeUrl string
}

// NewAlibabadropshippingstoresaveRequest 初始化AlibabadropshippingstoresaveAPIRequest对象
func NewAlibabadropshippingstoresaveRequest() *AlibabadropshippingstoresaveAPIRequest {
	return &AlibabadropshippingstoresaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadropshippingstoresaveAPIRequest) GetApiMethodName() string {
	return "alibaba.dropshipping.store.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadropshippingstoresaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadropshippingstoresaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreType is StoreType Setter
// store type
func (r *AlibabadropshippingstoresaveAPIRequest) SetStoreType(_storeType string) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// GetStoreType StoreType Getter
func (r AlibabadropshippingstoresaveAPIRequest) GetStoreType() string {
	return r._storeType
}

// SetStoreUrl is StoreUrl Setter
// store url
func (r *AlibabadropshippingstoresaveAPIRequest) SetStoreUrl(_storeUrl string) error {
	r._storeUrl = _storeUrl
	r.Set("store_url", _storeUrl)
	return nil
}

// GetStoreUrl StoreUrl Getter
func (r AlibabadropshippingstoresaveAPIRequest) GetStoreUrl() string {
	return r._storeUrl
}
