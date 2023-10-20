package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthDeviceGet 获取设备详情
// alibaba.ailabs.tmallgenie.auth.device.get
//
// 通过此接口获取设备详情
func AlibabaAilabsTmallgenieAuthDeviceGet(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
