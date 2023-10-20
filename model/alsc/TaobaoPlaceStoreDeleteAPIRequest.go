package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoredeleteAPIRequest 线下门店删除 API请求
// taobao.place.store.delete
//
// 用于商家删除线下门店
type TaobaoplacestoredeleteAPIRequest struct {
	model.Params
	// 门店id
	_storeId int64
}

// NewTaobaoplacestoredeleteRequest 初始化TaobaoplacestoredeleteAPIRequest对象
func NewTaobaoplacestoredeleteRequest() *TaobaoplacestoredeleteAPIRequest {
	return &TaobaoplacestoredeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestoredeleteAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestoredeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestoredeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店id
func (r *TaobaoplacestoredeleteAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoplacestoredeleteAPIRequest) GetStoreId() int64 {
	return r._storeId
}
