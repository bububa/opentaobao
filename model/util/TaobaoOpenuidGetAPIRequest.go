package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenuidGetAPIRequest 获取授权账号对应的OpenUid API请求
// taobao.openuid.get
//
// 获取授权账号对应的OpenUid
type TaobaoOpenuidGetAPIRequest struct {
	model.Params
}

// NewTaobaoOpenuidGetRequest 初始化TaobaoOpenuidGetAPIRequest对象
func NewTaobaoOpenuidGetRequest() *TaobaoOpenuidGetAPIRequest {
	return &TaobaoOpenuidGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenuidGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenuidGetAPIRequest) GetApiMethodName() string {
	return "taobao.openuid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenuidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenuidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoOpenuidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenuidGetRequest()
	},
}

// GetTaobaoOpenuidGetRequest 从 sync.Pool 获取 TaobaoOpenuidGetAPIRequest
func GetTaobaoOpenuidGetAPIRequest() *TaobaoOpenuidGetAPIRequest {
	return poolTaobaoOpenuidGetAPIRequest.Get().(*TaobaoOpenuidGetAPIRequest)
}

// ReleaseTaobaoOpenuidGetAPIRequest 将 TaobaoOpenuidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenuidGetAPIRequest(v *TaobaoOpenuidGetAPIRequest) {
	v.Reset()
	poolTaobaoOpenuidGetAPIRequest.Put(v)
}
