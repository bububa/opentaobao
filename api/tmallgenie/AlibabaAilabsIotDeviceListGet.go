package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsIotDeviceListGet 获取iot设备列表
// alibaba.ailabs.iot.device.list.get
//
// 通过此接口获取用户名下的iot设备列表
func AlibabaAilabsIotDeviceListGet(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsIotDeviceListGetAPIRequest, resp *tmallgenie.AlibabaAilabsIotDeviceListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
