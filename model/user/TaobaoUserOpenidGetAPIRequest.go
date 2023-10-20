package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserOpenidGetAPIRequest 查询用户openId API请求
// taobao.user.openid.get
//
// 查询用户openId
type TaobaoUserOpenidGetAPIRequest struct {
	model.Params
}

// NewTaobaoUserOpenidGetRequest 初始化TaobaoUserOpenidGetAPIRequest对象
func NewTaobaoUserOpenidGetRequest() *TaobaoUserOpenidGetAPIRequest {
	return &TaobaoUserOpenidGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUserOpenidGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserOpenidGetAPIRequest) GetApiMethodName() string {
	return "taobao.user.openid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserOpenidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUserOpenidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoUserOpenidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUserOpenidGetRequest()
	},
}

// GetTaobaoUserOpenidGetRequest 从 sync.Pool 获取 TaobaoUserOpenidGetAPIRequest
func GetTaobaoUserOpenidGetAPIRequest() *TaobaoUserOpenidGetAPIRequest {
	return poolTaobaoUserOpenidGetAPIRequest.Get().(*TaobaoUserOpenidGetAPIRequest)
}

// ReleaseTaobaoUserOpenidGetAPIRequest 将 TaobaoUserOpenidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUserOpenidGetAPIRequest(v *TaobaoUserOpenidGetAPIRequest) {
	v.Reset()
	poolTaobaoUserOpenidGetAPIRequest.Put(v)
}
