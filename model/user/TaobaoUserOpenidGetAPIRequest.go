package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouseropenidgetAPIRequest 查询用户openId API请求
// taobao.user.openid.get
//
// 查询用户openId
type TaobaouseropenidgetAPIRequest struct {
	model.Params
}

// NewTaobaouseropenidgetRequest 初始化TaobaouseropenidgetAPIRequest对象
func NewTaobaouseropenidgetRequest() *TaobaouseropenidgetAPIRequest {
	return &TaobaouseropenidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouseropenidgetAPIRequest) GetApiMethodName() string {
	return "taobao.user.openid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouseropenidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouseropenidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
