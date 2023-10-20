package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreModifyAPIRequest 商家修改线下门店 API请求
// taobao.place.store.modify
//
// 用于商家修改线下门店信息
type TaobaoPlaceStoreModifyAPIRequest struct {
	model.Params
	// 门店创建入参
	_storeUpdate *StoreUpdateTopDto
}

// NewTaobaoPlaceStoreModifyRequest 初始化TaobaoPlaceStoreModifyAPIRequest对象
func NewTaobaoPlaceStoreModifyRequest() *TaobaoPlaceStoreModifyAPIRequest {
	return &TaobaoPlaceStoreModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStoreModifyAPIRequest) Reset() {
	r._storeUpdate = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreModifyAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoreModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreUpdate is StoreUpdate Setter
// 门店创建入参
func (r *TaobaoPlaceStoreModifyAPIRequest) SetStoreUpdate(_storeUpdate *StoreUpdateTopDto) error {
	r._storeUpdate = _storeUpdate
	r.Set("store_update", _storeUpdate)
	return nil
}

// GetStoreUpdate StoreUpdate Getter
func (r TaobaoPlaceStoreModifyAPIRequest) GetStoreUpdate() *StoreUpdateTopDto {
	return r._storeUpdate
}

var poolTaobaoPlaceStoreModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStoreModifyRequest()
	},
}

// GetTaobaoPlaceStoreModifyRequest 从 sync.Pool 获取 TaobaoPlaceStoreModifyAPIRequest
func GetTaobaoPlaceStoreModifyAPIRequest() *TaobaoPlaceStoreModifyAPIRequest {
	return poolTaobaoPlaceStoreModifyAPIRequest.Get().(*TaobaoPlaceStoreModifyAPIRequest)
}

// ReleaseTaobaoPlaceStoreModifyAPIRequest 将 TaobaoPlaceStoreModifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStoreModifyAPIRequest(v *TaobaoPlaceStoreModifyAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStoreModifyAPIRequest.Put(v)
}
