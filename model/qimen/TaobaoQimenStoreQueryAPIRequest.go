package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreQueryAPIRequest 门店信息查询接口 API请求
// taobao.qimen.store.query
//
// 商家在ERP等系统中调用该接口，查询门店相关信息
type TaobaoQimenStoreQueryAPIRequest struct {
	model.Params
	// 已分配的线上门店ID
	_storeId int64
}

// NewTaobaoQimenStoreQueryRequest 初始化TaobaoQimenStoreQueryAPIRequest对象
func NewTaobaoQimenStoreQueryRequest() *TaobaoQimenStoreQueryAPIRequest {
	return &TaobaoQimenStoreQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenStoreQueryAPIRequest) Reset() {
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStoreQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 已分配的线上门店ID
func (r *TaobaoQimenStoreQueryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoQimenStoreQueryAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoQimenStoreQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenStoreQueryRequest()
	},
}

// GetTaobaoQimenStoreQueryRequest 从 sync.Pool 获取 TaobaoQimenStoreQueryAPIRequest
func GetTaobaoQimenStoreQueryAPIRequest() *TaobaoQimenStoreQueryAPIRequest {
	return poolTaobaoQimenStoreQueryAPIRequest.Get().(*TaobaoQimenStoreQueryAPIRequest)
}

// ReleaseTaobaoQimenStoreQueryAPIRequest 将 TaobaoQimenStoreQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenStoreQueryAPIRequest(v *TaobaoQimenStoreQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenStoreQueryAPIRequest.Put(v)
}
