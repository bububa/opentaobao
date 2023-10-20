package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorecreateAPIRequest 商户门店创建接口 API请求
// taobao.place.store.create
//
// 用于商家创建线下门店
type TaobaoplacestorecreateAPIRequest struct {
	model.Params
	// 门店创建入参
	_storeCreate *StoreUpdateTopDto
}

// NewTaobaoplacestorecreateRequest 初始化TaobaoplacestorecreateAPIRequest对象
func NewTaobaoplacestorecreateRequest() *TaobaoplacestorecreateAPIRequest {
	return &TaobaoplacestorecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestorecreateAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestorecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestorecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCreate is StoreCreate Setter
// 门店创建入参
func (r *TaobaoplacestorecreateAPIRequest) SetStoreCreate(_storeCreate *StoreUpdateTopDto) error {
	r._storeCreate = _storeCreate
	r.Set("store_create", _storeCreate)
	return nil
}

// GetStoreCreate StoreCreate Getter
func (r TaobaoplacestorecreateAPIRequest) GetStoreCreate() *StoreUpdateTopDto {
	return r._storeCreate
}
