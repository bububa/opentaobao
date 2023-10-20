package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCustomersSidGetAPIRequest 查看功能权限 API请求
// taobao.simba.customers.sid.get
//
// 查询用户是否拥有某个功能权限
type TaobaoSimbaCustomersSidGetAPIRequest struct {
	model.Params
}

// NewTaobaoSimbaCustomersSidGetRequest 初始化TaobaoSimbaCustomersSidGetAPIRequest对象
func NewTaobaoSimbaCustomersSidGetRequest() *TaobaoSimbaCustomersSidGetAPIRequest {
	return &TaobaoSimbaCustomersSidGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCustomersSidGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCustomersSidGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.customers.sid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCustomersSidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCustomersSidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoSimbaCustomersSidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCustomersSidGetRequest()
	},
}

// GetTaobaoSimbaCustomersSidGetRequest 从 sync.Pool 获取 TaobaoSimbaCustomersSidGetAPIRequest
func GetTaobaoSimbaCustomersSidGetAPIRequest() *TaobaoSimbaCustomersSidGetAPIRequest {
	return poolTaobaoSimbaCustomersSidGetAPIRequest.Get().(*TaobaoSimbaCustomersSidGetAPIRequest)
}

// ReleaseTaobaoSimbaCustomersSidGetAPIRequest 将 TaobaoSimbaCustomersSidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCustomersSidGetAPIRequest(v *TaobaoSimbaCustomersSidGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCustomersSidGetAPIRequest.Put(v)
}
