package yunosminiapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosMiniappDatatunnelCallAPIRequest
车载小程序外部服务调用 API请求
yunos.miniapp.datatunnel.call

对客户提供的api进行统一封装调用。 */
type YunosMiniappDatatunnelCallAPIRequest struct {
	model.Params
	// 参数
	_param *YunosMiniappDatatunnelCallBaseRequest
}

// NewYunosMiniappDatatunnelCallRequest 初始化YunosMiniappDatatunnelCallAPIRequest对象
func NewYunosMiniappDatatunnelCallRequest() *YunosMiniappDatatunnelCallAPIRequest {
	return &YunosMiniappDatatunnelCallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosMiniappDatatunnelCallAPIRequest) GetApiMethodName() string {
	return "yunos.miniapp.datatunnel.call"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosMiniappDatatunnelCallAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 参数
func (r *YunosMiniappDatatunnelCallAPIRequest) SetParam(_param *YunosMiniappDatatunnelCallBaseRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r YunosMiniappDatatunnelCallAPIRequest) GetParam() *YunosMiniappDatatunnelCallBaseRequest {
	return r._param
}
