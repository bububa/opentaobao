package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsIotDeviceControlNotify 天猫精灵IoT异步控制回调接口
// alibaba.ailabs.iot.device.control.notify
//
// 用于天猫精灵IoT云云接入控制结果的异步回调
func AlibabaAilabsIotDeviceControlNotify(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsIotDeviceControlNotifyAPIRequest, resp *tmallgenie.AlibabaAilabsIotDeviceControlNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
