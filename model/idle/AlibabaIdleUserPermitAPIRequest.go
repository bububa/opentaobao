package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleUserPermitAPIRequest 用户appkey授权 API请求
// alibaba.idle.user.permit
//
// 用于记录登录用户与服务商的绑定关系，用于业务数据分发和授权校验
type AlibabaIdleUserPermitAPIRequest struct {
	model.Params
	// 授权请求
	_paramUserGrantRequest *UserGrantRequest
}

// NewAlibabaIdleUserPermitRequest 初始化AlibabaIdleUserPermitAPIRequest对象
func NewAlibabaIdleUserPermitRequest() *AlibabaIdleUserPermitAPIRequest {
	return &AlibabaIdleUserPermitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleUserPermitAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleUserPermitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamUserGrantRequest is ParamUserGrantRequest Setter
// 授权请求
func (r *AlibabaIdleUserPermitAPIRequest) SetParamUserGrantRequest(_paramUserGrantRequest *UserGrantRequest) error {
	r._paramUserGrantRequest = _paramUserGrantRequest
	r.Set("param_user_grant_request", _paramUserGrantRequest)
	return nil
}

// GetParamUserGrantRequest ParamUserGrantRequest Getter
func (r AlibabaIdleUserPermitAPIRequest) GetParamUserGrantRequest() *UserGrantRequest {
	return r._paramUserGrantRequest
}
