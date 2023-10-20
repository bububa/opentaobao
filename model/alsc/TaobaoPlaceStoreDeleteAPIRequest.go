package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreDeleteAPIRequest 线下门店删除 API请求
// taobao.place.store.delete
//
// 用于商家删除线下门店
type TaobaoPlaceStoreDeleteAPIRequest struct {
	model.Params
	// 门店id
	_storeId int64
}

// NewTaobaoPlaceStoreDeleteRequest 初始化TaobaoPlaceStoreDeleteAPIRequest对象
func NewTaobaoPlaceStoreDeleteRequest() *TaobaoPlaceStoreDeleteAPIRequest {
	return &TaobaoPlaceStoreDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStoreDeleteAPIRequest) Reset() {
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoreDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店id
func (r *TaobaoPlaceStoreDeleteAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoPlaceStoreDeleteAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoPlaceStoreDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStoreDeleteRequest()
	},
}

// GetTaobaoPlaceStoreDeleteRequest 从 sync.Pool 获取 TaobaoPlaceStoreDeleteAPIRequest
func GetTaobaoPlaceStoreDeleteAPIRequest() *TaobaoPlaceStoreDeleteAPIRequest {
	return poolTaobaoPlaceStoreDeleteAPIRequest.Get().(*TaobaoPlaceStoreDeleteAPIRequest)
}

// ReleaseTaobaoPlaceStoreDeleteAPIRequest 将 TaobaoPlaceStoreDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStoreDeleteAPIRequest(v *TaobaoPlaceStoreDeleteAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStoreDeleteAPIRequest.Put(v)
}
