package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappuserInfogetAPIRequest 用户开放信息获取 API请求
// taobao.miniapp.userInfo.get
//
// 获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
type TaobaominiappuserInfogetAPIRequest struct {
	model.Params
}

// NewTaobaominiappuserInfogetRequest 初始化TaobaominiappuserInfogetAPIRequest对象
func NewTaobaominiappuserInfogetRequest() *TaobaominiappuserInfogetAPIRequest {
	return &TaobaominiappuserInfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappuserInfogetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.userInfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappuserInfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappuserInfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
