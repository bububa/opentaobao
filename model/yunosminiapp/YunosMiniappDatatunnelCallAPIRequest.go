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
	_param *BaseRequest
}

// New
