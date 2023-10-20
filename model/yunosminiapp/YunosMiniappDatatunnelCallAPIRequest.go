package yunosminiapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosminiappdatatunnelcallAPIRequest 车载小程序外部服务调用 API请求
// yunos.miniapp.datatunnel.call
//
// 对客户提供的api进行统一封装调用。
type YunosminiappdatatunnelcallAPIRequest struct {
	model.Params
	// 参数
	_param *YunosminiappdatatunnelcallBaseRequest
}

// NewYunosminiappdatatunnelcallRequest 初始化YunosminiappdatatunnelcallAPIRequest对象
func NewYunosminiappdatatunnelcallRequest() *YunosminiappdatatunnelcallAPIRequest {
	return &YunosminiappdatatunnelcallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosminiappdatatunnelcallAPIRequest) GetApiMethodName() string {
	return "yunos.miniapp.datatunnel.call"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosminiappdatatunnelcallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosminiappdatatunnelcallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数
func (r *YunosminiappdatatunnelcallAPIRequest) SetParam(_param *YunosminiappdatatunnelcallBaseRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r YunosminiappdatatunnelcallAPIRequest) GetParam() *YunosminiappdatatunnelcallBaseRequest {
	return r._param
}
