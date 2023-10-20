package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaologinusergetAPIRequest 获取分销用户登录信息 API请求
// taobao.fenxiao.login.user.get
//
// 获取用户登录信息
type TaobaofenxiaologinusergetAPIRequest struct {
	model.Params
}

// NewTaobaofenxiaologinusergetRequest 初始化TaobaofenxiaologinusergetAPIRequest对象
func NewTaobaofenxiaologinusergetRequest() *TaobaofenxiaologinusergetAPIRequest {
	return &TaobaofenxiaologinusergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaologinusergetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.login.user.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaologinusergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaologinusergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
