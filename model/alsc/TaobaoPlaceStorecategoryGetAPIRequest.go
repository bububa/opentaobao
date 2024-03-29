package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStorecategoryGetAPIRequest 获取门店类目信息 API请求
// taobao.place.storecategory.get
//
// 获取门店类目信息
type TaobaoPlaceStorecategoryGetAPIRequest struct {
	model.Params
}

// NewTaobaoPlaceStorecategoryGetRequest 初始化TaobaoPlaceStorecategoryGetAPIRequest对象
func NewTaobaoPlaceStorecategoryGetRequest() *TaobaoPlaceStorecategoryGetAPIRequest {
	return &TaobaoPlaceStorecategoryGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStorecategoryGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorecategoryGetAPIRequest) GetApiMethodName() string {
	return "taobao.place.storecategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorecategoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStorecategoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoPlaceStorecategoryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStorecategoryGetRequest()
	},
}

// GetTaobaoPlaceStorecategoryGetRequest 从 sync.Pool 获取 TaobaoPlaceStorecategoryGetAPIRequest
func GetTaobaoPlaceStorecategoryGetAPIRequest() *TaobaoPlaceStorecategoryGetAPIRequest {
	return poolTaobaoPlaceStorecategoryGetAPIRequest.Get().(*TaobaoPlaceStorecategoryGetAPIRequest)
}

// ReleaseTaobaoPlaceStorecategoryGetAPIRequest 将 TaobaoPlaceStorecategoryGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStorecategoryGetAPIRequest(v *TaobaoPlaceStorecategoryGetAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStorecategoryGetAPIRequest.Put(v)
}
