package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalitemunbindAPIRequest ISV解绑商品 API请求
// alibaba.alihealth.dental.item.unbind
//
// ISV解绑商品
type AlibabaalihealthdentalitemunbindAPIRequest struct {
	model.Params
	// ISV门店ID
	_spStoreId string
	// ISV商品ID
	_spItemId string
}

// NewAlibabaalihealthdentalitemunbindRequest 初始化AlibabaalihealthdentalitemunbindAPIRequest对象
func NewAlibabaalihealthdentalitemunbindRequest() *AlibabaalihealthdentalitemunbindAPIRequest {
	return &AlibabaalihealthdentalitemunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdentalitemunbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.item.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdentalitemunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdentalitemunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpStoreId is SpStoreId Setter
// ISV门店ID
func (r *AlibabaalihealthdentalitemunbindAPIRequest) SetSpStoreId(_spStoreId string) error {
	r._spStoreId = _spStoreId
	r.Set("sp_store_id", _spStoreId)
	return nil
}

// GetSpStoreId SpStoreId Getter
func (r AlibabaalihealthdentalitemunbindAPIRequest) GetSpStoreId() string {
	return r._spStoreId
}

// SetSpItemId is SpItemId Setter
// ISV商品ID
func (r *AlibabaalihealthdentalitemunbindAPIRequest) SetSpItemId(_spItemId string) error {
	r._spItemId = _spItemId
	r.Set("sp_item_id", _spItemId)
	return nil
}

// GetSpItemId SpItemId Getter
func (r AlibabaalihealthdentalitemunbindAPIRequest) GetSpItemId() string {
	return r._spItemId
}
