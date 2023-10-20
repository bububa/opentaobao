package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleuserpermitAPIRequest 用户appkey授权 API请求
// alibaba.idle.user.permit
//
// 闲鱼卖家与服务商关系绑定，用于业务数据分发/授权校验/消息通知鉴权
type AlibabaidleuserpermitAPIRequest struct {
	model.Params
	// 授权请求
	_paramUserGrantRequest *UserGrantRequest
}

// NewAlibabaidleuserpermitRequest 初始化AlibabaidleuserpermitAPIRequest对象
func NewAlibabaidleuserpermitRequest() *AlibabaidleuserpermitAPIRequest {
	return &AlibabaidleuserpermitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleuserpermitAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleuserpermitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleuserpermitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUserGrantRequest is ParamUserGrantRequest Setter
// 授权请求
func (r *AlibabaidleuserpermitAPIRequest) SetParamUserGrantRequest(_paramUserGrantRequest *UserGrantRequest) error {
	r._paramUserGrantRequest = _paramUserGrantRequest
	r.Set("param_user_grant_request", _paramUserGrantRequest)
	return nil
}

// GetParamUserGrantRequest ParamUserGrantRequest Getter
func (r AlibabaidleuserpermitAPIRequest) GetParamUserGrantRequest() *UserGrantRequest {
	return r._paramUserGrantRequest
}
