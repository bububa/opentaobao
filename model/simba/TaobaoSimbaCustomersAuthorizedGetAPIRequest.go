package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacustomersauthorizedgetAPIRequest 取得当前登录用户的授权账户列表 API请求
// taobao.simba.customers.authorized.get
//
// 取得当前登录用户的授权账户列表
type TaobaosimbacustomersauthorizedgetAPIRequest struct {
	model.Params
}

// NewTaobaosimbacustomersauthorizedgetRequest 初始化TaobaosimbacustomersauthorizedgetAPIRequest对象
func NewTaobaosimbacustomersauthorizedgetRequest() *TaobaosimbacustomersauthorizedgetAPIRequest {
	return &TaobaosimbacustomersauthorizedgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacustomersauthorizedgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.customers.authorized.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacustomersauthorizedgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacustomersauthorizedgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
