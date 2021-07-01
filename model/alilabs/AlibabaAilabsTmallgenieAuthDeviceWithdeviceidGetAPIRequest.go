package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest
根据三方ID查询设备注册激活信息 API请求
alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get

根据三方ID查询设备注册激活信息 */
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest struct {
	model.Params
	// 设备产品ID
	_clientId string
	// mac地址
	_mac string
}

// New
