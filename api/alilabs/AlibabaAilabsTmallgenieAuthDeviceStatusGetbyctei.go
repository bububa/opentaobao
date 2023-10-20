package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGetbyctei 根据ctei查询状态
// alibaba.ailabs.tmallgenie.auth.device.status.getbyctei
//
// 提供给电信查询设备在线状态值
func AlibabaAilabsTmallgenieAuthDeviceStatusGetbyctei(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
