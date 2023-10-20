package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappeleuserphonegetAPIRequest 获取饿了么用户信息 API请求
// taobao.miniapp.eleuser.phone.get
//
// 获取饿了么用户信息
type TaobaominiappeleuserphonegetAPIRequest struct {
	model.Params
}

// NewTaobaominiappeleuserphonegetRequest 初始化TaobaominiappeleuserphonegetAPIRequest对象
func NewTaobaominiappeleuserphonegetRequest() *TaobaominiappeleuserphonegetAPIRequest {
	return &TaobaominiappeleuserphonegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappeleuserphonegetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.eleuser.phone.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappeleuserphonegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappeleuserphonegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
