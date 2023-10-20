package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstoredeleteAPIRequest 门店删除接口 API请求
// taobao.qimen.store.delete
//
// 商家在ERP等系统中调用该接口，删除线下门店
type TaobaoqimenstoredeleteAPIRequest struct {
	model.Params
	// 要删除的门店id
	_storeId int64
}

// NewTaobaoqimenstoredeleteRequest 初始化TaobaoqimenstoredeleteAPIRequest对象
func NewTaobaoqimenstoredeleteRequest() *TaobaoqimenstoredeleteAPIRequest {
	return &TaobaoqimenstoredeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstoredeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstoredeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstoredeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 要删除的门店id
func (r *TaobaoqimenstoredeleteAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoqimenstoredeleteAPIRequest) GetStoreId() int64 {
	return r._storeId
}
