package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCustomersAuthorizedGetAPIRequest 取得当前登录用户的授权账户列表 API请求
// taobao.simba.customers.authorized.get
//
// 取得当前登录用户的授权账户列表
type TaobaoSimbaCustomersAuthorizedGetAPIRequest struct {
	model.Params
}

// NewTaobaoSimbaCustomersAuthorizedGetRequest 初始化TaobaoSimbaCustomersAuthorizedGetAPIRequest对象
func NewTaobaoSimbaCustomersAuthorizedGetRequest() *TaobaoSimbaCustomersAuthorizedGetAPIRequest {
	return &TaobaoSimbaCustomersAuthorizedGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCustomersAuthorizedGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCustomersAuthorizedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.customers.authorized.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCustomersAuthorizedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCustomersAuthorizedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoSimbaCustomersAuthorizedGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCustomersAuthorizedGetRequest()
	},
}

// GetTaobaoSimbaCustomersAuthorizedGetRequest 从 sync.Pool 获取 TaobaoSimbaCustomersAuthorizedGetAPIRequest
func GetTaobaoSimbaCustomersAuthorizedGetAPIRequest() *TaobaoSimbaCustomersAuthorizedGetAPIRequest {
	return poolTaobaoSimbaCustomersAuthorizedGetAPIRequest.Get().(*TaobaoSimbaCustomersAuthorizedGetAPIRequest)
}

// ReleaseTaobaoSimbaCustomersAuthorizedGetAPIRequest 将 TaobaoSimbaCustomersAuthorizedGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCustomersAuthorizedGetAPIRequest(v *TaobaoSimbaCustomersAuthorizedGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCustomersAuthorizedGetAPIRequest.Put(v)
}
