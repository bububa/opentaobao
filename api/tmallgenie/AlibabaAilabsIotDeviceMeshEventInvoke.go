package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsIotDeviceMeshEventInvoke 弹内设备中心事件调用
// alibaba.ailabs.iot.device.mesh.event.invoke
//
// 弹内设备中心事件调用
func AlibabaAilabsIotDeviceMeshEventInvoke(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest, session string) (*tmallgenie.AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
