package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacustomerssidgetAPIRequest 查看功能权限 API请求
// taobao.simba.customers.sid.get
//
// 查询用户是否拥有某个功能权限
type TaobaosimbacustomerssidgetAPIRequest struct {
	model.Params
}

// NewTaobaosimbacustomerssidgetRequest 初始化TaobaosimbacustomerssidgetAPIRequest对象
func NewTaobaosimbacustomerssidgetRequest() *TaobaosimbacustomerssidgetAPIRequest {
	return &TaobaosimbacustomerssidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacustomerssidgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.customers.sid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacustomerssidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacustomerssidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
