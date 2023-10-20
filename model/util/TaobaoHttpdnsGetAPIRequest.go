package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoHttpdnsGetAPIRequest TOPDNS配置 API请求
// taobao.httpdns.get
//
// 获取TOP DNS配置
type TaobaoHttpdnsGetAPIRequest struct {
	model.Params
}

// NewTaobaoHttpdnsGetRequest 初始化TaobaoHttpdnsGetAPIRequest对象
func NewTaobaoHttpdnsGetRequest() *TaobaoHttpdnsGetAPIRequest {
	return &TaobaoHttpdnsGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoHttpdnsGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoHttpdnsGetAPIRequest) GetApiMethodName() string {
	return "taobao.httpdns.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoHttpdnsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoHttpdnsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoHttpdnsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoHttpdnsGetRequest()
	},
}

// GetTaobaoHttpdnsGetRequest 从 sync.Pool 获取 TaobaoHttpdnsGetAPIRequest
func GetTaobaoHttpdnsGetAPIRequest() *TaobaoHttpdnsGetAPIRequest {
	return poolTaobaoHttpdnsGetAPIRequest.Get().(*TaobaoHttpdnsGetAPIRequest)
}

// ReleaseTaobaoHttpdnsGetAPIRequest 将 TaobaoHttpdnsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoHttpdnsGetAPIRequest(v *TaobaoHttpdnsGetAPIRequest) {
	v.Reset()
	poolTaobaoHttpdnsGetAPIRequest.Put(v)
}
