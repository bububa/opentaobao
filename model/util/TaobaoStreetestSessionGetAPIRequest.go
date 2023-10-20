package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoStreetestSessionGetAPIRequest 根据获取压测用户的sessionKey API请求
// taobao.streetest.session.get
//
// 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
type TaobaoStreetestSessionGetAPIRequest struct {
	model.Params
}

// NewTaobaoStreetestSessionGetRequest 初始化TaobaoStreetestSessionGetAPIRequest对象
func NewTaobaoStreetestSessionGetRequest() *TaobaoStreetestSessionGetAPIRequest {
	return &TaobaoStreetestSessionGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoStreetestSessionGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoStreetestSessionGetAPIRequest) GetApiMethodName() string {
	return "taobao.streetest.session.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoStreetestSessionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoStreetestSessionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoStreetestSessionGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoStreetestSessionGetRequest()
	},
}

// GetTaobaoStreetestSessionGetRequest 从 sync.Pool 获取 TaobaoStreetestSessionGetAPIRequest
func GetTaobaoStreetestSessionGetAPIRequest() *TaobaoStreetestSessionGetAPIRequest {
	return poolTaobaoStreetestSessionGetAPIRequest.Get().(*TaobaoStreetestSessionGetAPIRequest)
}

// ReleaseTaobaoStreetestSessionGetAPIRequest 将 TaobaoStreetestSessionGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoStreetestSessionGetAPIRequest(v *TaobaoStreetestSessionGetAPIRequest) {
	v.Reset()
	poolTaobaoStreetestSessionGetAPIRequest.Put(v)
}
