package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovasservicevalidateAPIRequest 增值服务订购服务验证 API请求
// taobao.vas.service.validate
//
// 增值服务订购服务验证
type TaobaovasservicevalidateAPIRequest struct {
	model.Params
	// 用户昵称
	_nick string
	// 服务编码
	_servCode string
}

// NewTaobaovasservicevalidateRequest 初始化TaobaovasservicevalidateAPIRequest对象
func NewTaobaovasservicevalidateRequest() *TaobaovasservicevalidateAPIRequest {
	return &TaobaovasservicevalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovasservicevalidateAPIRequest) GetApiMethodName() string {
	return "taobao.vas.service.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovasservicevalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovasservicevalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaovasservicevalidateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaovasservicevalidateAPIRequest) GetNick() string {
	return r._nick
}

// SetServCode is ServCode Setter
// 服务编码
func (r *TaobaovasservicevalidateAPIRequest) SetServCode(_servCode string) error {
	r._servCode = _servCode
	r.Set("serv_code", _servCode)
	return nil
}

// GetServCode ServCode Getter
func (r TaobaovasservicevalidateAPIRequest) GetServCode() string {
	return r._servCode
}
