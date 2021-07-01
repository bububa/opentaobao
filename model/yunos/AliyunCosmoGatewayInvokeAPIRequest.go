package yunos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunCosmoGatewayInvokeAPIRequest
alios cosmo服务调用 API请求
aliyun.cosmo.gateway.invoke

AliOS cosmo服务分发平台对外调用接口 */
type AliyunCosmoGatewayInvokeAPIRequest struct {
	model.Params
	// 请求上下文参数
	_context *RdamContext
	// 请求对象
	_rdamRequest *RdamGenericRequest
}

// New
