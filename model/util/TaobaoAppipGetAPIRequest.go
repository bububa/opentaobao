package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppipGetAPIRequest 获取ISV发起请求服务器IP API请求
// taobao.appip.get
//
// 获取ISV发起请求服务器IP
type TaobaoAppipGetAPIRequest struct {
	model.Params
}

// NewTaobaoAppipGetRequest 初始化TaobaoAppipGetAPIRequest对象
func NewTaobaoAppipGetRequest() *TaobaoAppipGetAPIRequest {
	return &TaobaoAppipGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAppipGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppipGetAPIRequest) GetApiMethodName() string {
	return "taobao.appip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppipGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAppipGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoAppipGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAppipGetRequest()
	},
}

// GetTaobaoAppipGetRequest 从 sync.Pool 获取 TaobaoAppipGetAPIRequest
func GetTaobaoAppipGetAPIRequest() *TaobaoAppipGetAPIRequest {
	return poolTaobaoAppipGetAPIRequest.Get().(*TaobaoAppipGetAPIRequest)
}

// ReleaseTaobaoAppipGetAPIRequest 将 TaobaoAppipGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAppipGetAPIRequest(v *TaobaoAppipGetAPIRequest) {
	v.Reset()
	poolTaobaoAppipGetAPIRequest.Put(v)
}
