package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoWtUserCrowdAPIRequest 校验用户是否为新人 API请求
// alibaba.taobao.wt.user.crowd
//
// 增加isv接口
// 根据入参手机号判断是否为新人类型
type AlibabaTaobaoWtUserCrowdAPIRequest struct {
	model.Params
	// 手机号
	_phone int64
}

// NewAlibabaTaobaoWtUserCrowdRequest 初始化AlibabaTaobaoWtUserCrowdAPIRequest对象
func NewAlibabaTaobaoWtUserCrowdRequest() *AlibabaTaobaoWtUserCrowdAPIRequest {
	return &AlibabaTaobaoWtUserCrowdAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoWtUserCrowdAPIRequest) GetApiMethodName() string {
	return "alibaba.taobao.wt.user.crowd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoWtUserCrowdAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Phone Setter
// 手机号
func (r *AlibabaTaobaoWtUserCrowdAPIRequest) SetPhone(_phone int64) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// Get Phone Getter
func (r AlibabaTaobaoWtUserCrowdAPIRequest) GetPhone() int64 {
	return r._phone
}
