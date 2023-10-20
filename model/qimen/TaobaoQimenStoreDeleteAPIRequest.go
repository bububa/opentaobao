package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreDeleteAPIRequest 门店删除接口 API请求
// taobao.qimen.store.delete
//
// 商家在ERP等系统中调用该接口，删除线下门店
type TaobaoQimenStoreDeleteAPIRequest struct {
	model.Params
	// 要删除的门店id
	_storeId int64
}

// NewTaobaoQimenStoreDeleteRequest 初始化TaobaoQimenStoreDeleteAPIRequest对象
func NewTaobaoQimenStoreDeleteRequest() *TaobaoQimenStoreDeleteAPIRequest {
	return &TaobaoQimenStoreDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenStoreDeleteAPIRequest) Reset() {
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStoreDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 要删除的门店id
func (r *TaobaoQimenStoreDeleteAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoQimenStoreDeleteAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoQimenStoreDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenStoreDeleteRequest()
	},
}

// GetTaobaoQimenStoreDeleteRequest 从 sync.Pool 获取 TaobaoQimenStoreDeleteAPIRequest
func GetTaobaoQimenStoreDeleteAPIRequest() *TaobaoQimenStoreDeleteAPIRequest {
	return poolTaobaoQimenStoreDeleteAPIRequest.Get().(*TaobaoQimenStoreDeleteAPIRequest)
}

// ReleaseTaobaoQimenStoreDeleteAPIRequest 将 TaobaoQimenStoreDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenStoreDeleteAPIRequest(v *TaobaoQimenStoreDeleteAPIRequest) {
	v.Reset()
	poolTaobaoQimenStoreDeleteAPIRequest.Put(v)
}
