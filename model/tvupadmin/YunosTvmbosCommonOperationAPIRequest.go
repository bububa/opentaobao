package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvmbosCommonOperationAPIRequest
应用中心通用服务接口 API请求
yunos.tvmbos.common.operation

应用中心相关接口的代理 */
type YunosTvmbosCommonOperationAPIRequest struct {
	model.Params
	// 接口名
	_interfaceName string
	// 方法名
	_methodName string
	// 入参，json格式
	_parameter string
}

// New
