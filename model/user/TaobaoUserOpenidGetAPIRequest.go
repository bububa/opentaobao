package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserOpenidGetAPIRequest 查询用户openId API请求
// taobao.user.openid.get
//
// 查询用户openId
type TaobaoUserOpenidGetAPIRequest struct {
	model.Params
}

// NewTaobaoUserOpenidGetRequest 初始化TaobaoUserOpenidGetAPIRequest对象
func NewTaobaoUserOpenidGetRequest() *TaobaoUserOpenidGetAPIRequest {
	return &TaobaoUserOpenidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserOpenidGetAPIRequest) GetApiMethodName() string {
	return "taobao.user.openid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserOpenidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUserOpenidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
