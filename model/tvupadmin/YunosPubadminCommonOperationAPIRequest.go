package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosPubadminCommonOperationAPIRequest
内部迎客松通用服务 API请求
yunos.pubadmin.common.operation

内部迎客松通用服务 */
type YunosPubadminCommonOperationAPIRequest struct {
	model.Params
	// 入参json串
	_parameter string
	// 接口名
	_interfaceName string
	// 方法名
	_methodName string
}

// New
