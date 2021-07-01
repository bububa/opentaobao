package aligenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieDeviceUnbindAPIRequest
设备解绑操作接口 API请求
alibaba.ailabs.aligenie.device.unbind

让开发者能根据设备ID进行解绑操作的接口 */
type AlibabaAilabsAligenieDeviceUnbindAPIRequest struct {
	model.Params
	// 扩展信息，用于存放APP类型等
	_ext string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 账户体系隔离字符串
	_schema string
	// 欲解绑的设备ID
	_uuid string
}

// New
