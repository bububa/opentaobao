package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCustomersAuthorizedGetAPIRequest
取得当前登录用户的授权账户列表 API请求
taobao.simba.customers.authorized.get

取得当前登录用户的授权账户列表 */
type TaobaoSimbaCustomersAuthorizedGetAPIRequest struct {
	model.Params
}

// NewTaobaoSimbaCustomersAuthorizedGetRequest 初始化TaobaoSimbaCustomersAuthorizedGetAPIRequest对象
func NewTaobaoSimbaCustomersAuthorizedGetRequest() *TaobaoSimbaCustomersAuthorizedGetAPIRequest {
	return &TaobaoSimbaCustomersAuthorizedGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCustomersAuthorizedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.customers.authorized.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCustomersAuthorizedGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
