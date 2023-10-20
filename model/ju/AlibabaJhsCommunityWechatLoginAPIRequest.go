package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajhscommunitywechatloginAPIRequest 聚划算用增淘外社群登录 API请求
// alibaba.jhs.community.wechat.login
//
// 聚划算用增淘外社群登录
type AlibabajhscommunitywechatloginAPIRequest struct {
	model.Params
	// 动态令牌
	_code string
}

// NewAlibabajhscommunitywechatloginRequest 初始化AlibabajhscommunitywechatloginAPIRequest对象
func NewAlibabajhscommunitywechatloginRequest() *AlibabajhscommunitywechatloginAPIRequest {
	return &AlibabajhscommunitywechatloginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajhscommunitywechatloginAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.wechat.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajhscommunitywechatloginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajhscommunitywechatloginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 动态令牌
func (r *AlibabajhscommunitywechatloginAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabajhscommunitywechatloginAPIRequest) GetCode() string {
	return r._code
}
