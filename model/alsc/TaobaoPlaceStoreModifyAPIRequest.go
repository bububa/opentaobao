package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoremodifyAPIRequest 商家修改线下门店 API请求
// taobao.place.store.modify
//
// 用于商家修改线下门店信息
type TaobaoplacestoremodifyAPIRequest struct {
	model.Params
	// 门店创建入参
	_storeUpdate *StoreUpdateTopDto
}

// NewTaobaoplacestoremodifyRequest 初始化TaobaoplacestoremodifyAPIRequest对象
func NewTaobaoplacestoremodifyRequest() *TaobaoplacestoremodifyAPIRequest {
	return &TaobaoplacestoremodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestoremodifyAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestoremodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestoremodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreUpdate is StoreUpdate Setter
// 门店创建入参
func (r *TaobaoplacestoremodifyAPIRequest) SetStoreUpdate(_storeUpdate *StoreUpdateTopDto) error {
	r._storeUpdate = _storeUpdate
	r.Set("store_update", _storeUpdate)
	return nil
}

// GetStoreUpdate StoreUpdate Getter
func (r TaobaoplacestoremodifyAPIRequest) GetStoreUpdate() *StoreUpdateTopDto {
	return r._storeUpdate
}
