package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsHluserGetAPIRequest 订单全链路用户信息获取 API请求
// taobao.jds.hluser.get
//
// 订单全链路用户信息获取
type TaobaoJdsHluserGetAPIRequest struct {
	model.Params
}

// NewTaobaoJdsHluserGetRequest 初始化TaobaoJdsHluserGetAPIRequest对象
func NewTaobaoJdsHluserGetRequest() *TaobaoJdsHluserGetAPIRequest {
	return &TaobaoJdsHluserGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJdsHluserGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJdsHluserGetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.hluser.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJdsHluserGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJdsHluserGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoJdsHluserGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJdsHluserGetRequest()
	},
}

// GetTaobaoJdsHluserGetRequest 从 sync.Pool 获取 TaobaoJdsHluserGetAPIRequest
func GetTaobaoJdsHluserGetAPIRequest() *TaobaoJdsHluserGetAPIRequest {
	return poolTaobaoJdsHluserGetAPIRequest.Get().(*TaobaoJdsHluserGetAPIRequest)
}

// ReleaseTaobaoJdsHluserGetAPIRequest 将 TaobaoJdsHluserGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoJdsHluserGetAPIRequest(v *TaobaoJdsHluserGetAPIRequest) {
	v.Reset()
	poolTaobaoJdsHluserGetAPIRequest.Put(v)
}
