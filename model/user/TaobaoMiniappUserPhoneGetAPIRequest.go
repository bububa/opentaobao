package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappuserphonegetAPIRequest 获取当前授权用户手机号码 API请求
// taobao.miniapp.user.phone.get
//
// 在商家应用中，获取当前授权用户手机号码
type TaobaominiappuserphonegetAPIRequest struct {
	model.Params
}

// NewTaobaominiappuserphonegetRequest 初始化TaobaominiappuserphonegetAPIRequest对象
func NewTaobaominiappuserphonegetRequest() *TaobaominiappuserphonegetAPIRequest {
	return &TaobaominiappuserphonegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappuserphonegetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.user.phone.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappuserphonegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappuserphonegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
