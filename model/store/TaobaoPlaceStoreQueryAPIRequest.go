package store

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreQueryAPIRequest 门店信息查询接口 API请求
// taobao.place.store.query
//
// 根据用户授权信息，获取用户的门店公开信息
type TaobaoPlaceStoreQueryAPIRequest struct {
	model.Params
	// 业务code，用于区分业务
	_bizCode string
	// 业务外部id
	_outerId string
	// 门店id
	_storeId int64
}

// NewTaobaoPlaceStoreQueryRequest 初始化TaobaoPlaceStoreQueryAPIRequest对象
func NewTaobaoPlaceStoreQueryRequest() *TaobaoPlaceStoreQueryAPIRequest {
	return &TaobaoPlaceStoreQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStoreQueryAPIRequest) Reset() {
	r._bizCode = ""
	r._outerId = ""
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreQueryAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoreQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务code，用于区分业务
func (r *TaobaoPlaceStoreQueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TaobaoPlaceStoreQueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetOuterId is OuterId Setter
// 业务外部id
func (r *TaobaoPlaceStoreQueryAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoPlaceStoreQueryAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetStoreId is StoreId Setter
// 门店id
func (r *TaobaoPlaceStoreQueryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoPlaceStoreQueryAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoPlaceStoreQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStoreQueryRequest()
	},
}

// GetTaobaoPlaceStoreQueryRequest 从 sync.Pool 获取 TaobaoPlaceStoreQueryAPIRequest
func GetTaobaoPlaceStoreQueryAPIRequest() *TaobaoPlaceStoreQueryAPIRequest {
	return poolTaobaoPlaceStoreQueryAPIRequest.Get().(*TaobaoPlaceStoreQueryAPIRequest)
}

// ReleaseTaobaoPlaceStoreQueryAPIRequest 将 TaobaoPlaceStoreQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStoreQueryAPIRequest(v *TaobaoPlaceStoreQueryAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStoreQueryAPIRequest.Put(v)
}
