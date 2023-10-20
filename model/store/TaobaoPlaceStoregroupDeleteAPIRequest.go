package store

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoregroupDeleteAPIRequest 删除门店库 API请求
// taobao.place.storegroup.delete
//
// 删除门店库
type TaobaoPlaceStoregroupDeleteAPIRequest struct {
	model.Params
	// 库Id
	_id int64
}

// NewTaobaoPlaceStoregroupDeleteRequest 初始化TaobaoPlaceStoregroupDeleteAPIRequest对象
func NewTaobaoPlaceStoregroupDeleteRequest() *TaobaoPlaceStoregroupDeleteAPIRequest {
	return &TaobaoPlaceStoregroupDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStoregroupDeleteAPIRequest) Reset() {
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoregroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.place.storegroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoregroupDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoregroupDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 库Id
func (r *TaobaoPlaceStoregroupDeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoPlaceStoregroupDeleteAPIRequest) GetId() int64 {
	return r._id
}

var poolTaobaoPlaceStoregroupDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStoregroupDeleteRequest()
	},
}

// GetTaobaoPlaceStoregroupDeleteRequest 从 sync.Pool 获取 TaobaoPlaceStoregroupDeleteAPIRequest
func GetTaobaoPlaceStoregroupDeleteAPIRequest() *TaobaoPlaceStoregroupDeleteAPIRequest {
	return poolTaobaoPlaceStoregroupDeleteAPIRequest.Get().(*TaobaoPlaceStoregroupDeleteAPIRequest)
}

// ReleaseTaobaoPlaceStoregroupDeleteAPIRequest 将 TaobaoPlaceStoregroupDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStoregroupDeleteAPIRequest(v *TaobaoPlaceStoregroupDeleteAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStoregroupDeleteAPIRequest.Put(v)
}
