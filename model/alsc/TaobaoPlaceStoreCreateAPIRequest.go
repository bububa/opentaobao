package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreCreateAPIRequest 商户门店创建接口 API请求
// taobao.place.store.create
//
// 用于商家创建线下门店
type TaobaoPlaceStoreCreateAPIRequest struct {
	model.Params
	// 门店创建入参
	_storeCreate *StoreUpdateTopDto
}

// NewTaobaoPlaceStoreCreateRequest 初始化TaobaoPlaceStoreCreateAPIRequest对象
func NewTaobaoPlaceStoreCreateRequest() *TaobaoPlaceStoreCreateAPIRequest {
	return &TaobaoPlaceStoreCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStoreCreateAPIRequest) Reset() {
	r._storeCreate = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreCreateAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoreCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCreate is StoreCreate Setter
// 门店创建入参
func (r *TaobaoPlaceStoreCreateAPIRequest) SetStoreCreate(_storeCreate *StoreUpdateTopDto) error {
	r._storeCreate = _storeCreate
	r.Set("store_create", _storeCreate)
	return nil
}

// GetStoreCreate StoreCreate Getter
func (r TaobaoPlaceStoreCreateAPIRequest) GetStoreCreate() *StoreUpdateTopDto {
	return r._storeCreate
}

var poolTaobaoPlaceStoreCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStoreCreateRequest()
	},
}

// GetTaobaoPlaceStoreCreateRequest 从 sync.Pool 获取 TaobaoPlaceStoreCreateAPIRequest
func GetTaobaoPlaceStoreCreateAPIRequest() *TaobaoPlaceStoreCreateAPIRequest {
	return poolTaobaoPlaceStoreCreateAPIRequest.Get().(*TaobaoPlaceStoreCreateAPIRequest)
}

// ReleaseTaobaoPlaceStoreCreateAPIRequest 将 TaobaoPlaceStoreCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStoreCreateAPIRequest(v *TaobaoPlaceStoreCreateAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStoreCreateAPIRequest.Put(v)
}
