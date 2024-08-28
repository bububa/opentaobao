package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsIotDeviceMeshEventInvoke 弹内设备中心事件调用
// alibaba.ailabs.iot.device.mesh.event.invoke
//
// 弹内设备中心事件调用
func AlibabaAilabsIotDeviceMeshEventInvoke(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsIotDeviceMeshEventInvokeAPIRequest, resp *tmallgenie.AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
