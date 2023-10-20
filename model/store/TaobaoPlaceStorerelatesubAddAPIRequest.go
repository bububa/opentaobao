package store

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStorerelatesubAddAPIRequest 门店和子门店关系新增 API请求
// taobao.place.storerelatesub.add
//
// 门店和子门店关系新增
type TaobaoPlaceStorerelatesubAddAPIRequest struct {
	model.Params
	// 子门店Id
	_subStoreIds []string
	// 门店Id
	_storeId int64
}

// NewTaobaoPlaceStorerelatesubAddRequest 初始化TaobaoPlaceStorerelatesubAddAPIRequest对象
func NewTaobaoPlaceStorerelatesubAddRequest() *TaobaoPlaceStorerelatesubAddAPIRequest {
	return &TaobaoPlaceStorerelatesubAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStorerelatesubAddAPIRequest) Reset() {
	r._subStoreIds = r._subStoreIds[:0]
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetApiMethodName() string {
	return "taobao.place.storerelatesub.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubStoreIds is SubStoreIds Setter
// 子门店Id
func (r *TaobaoPlaceStorerelatesubAddAPIRequest) SetSubStoreIds(_subStoreIds []string) error {
	r._subStoreIds = _subStoreIds
	r.Set("sub_store_ids", _subStoreIds)
	return nil
}

// GetSubStoreIds SubStoreIds Getter
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetSubStoreIds() []string {
	return r._subStoreIds
}

// SetStoreId is StoreId Setter
// 门店Id
func (r *TaobaoPlaceStorerelatesubAddAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoPlaceStorerelatesubAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStorerelatesubAddRequest()
	},
}

// GetTaobaoPlaceStorerelatesubAddRequest 从 sync.Pool 获取 TaobaoPlaceStorerelatesubAddAPIRequest
func GetTaobaoPlaceStorerelatesubAddAPIRequest() *TaobaoPlaceStorerelatesubAddAPIRequest {
	return poolTaobaoPlaceStorerelatesubAddAPIRequest.Get().(*TaobaoPlaceStorerelatesubAddAPIRequest)
}

// ReleaseTaobaoPlaceStorerelatesubAddAPIRequest 将 TaobaoPlaceStorerelatesubAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStorerelatesubAddAPIRequest(v *TaobaoPlaceStorerelatesubAddAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStorerelatesubAddAPIRequest.Put(v)
}
