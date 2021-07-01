package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvscreenAdminCommonOperationAPIRequest
一体机桌面通用接口 API请求
yunos.tvscreen.admin.common.operation

一体机桌面通用接口 */
type YunosTvscreenAdminCommonOperationAPIRequest struct {
	model.Params
	// 参数数组
	_parameters string
	// 方法名
	_methodName string
	// 接口名称
	_interfaceName string
}

// New
