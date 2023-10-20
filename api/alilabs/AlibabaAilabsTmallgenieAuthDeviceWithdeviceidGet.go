package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGet 根据三方ID查询设备注册激活信息
// alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get
//
// 根据三方ID查询设备注册激活信息
func AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGet(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
