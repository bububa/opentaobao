package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotDeviceIsexist 判断设备是否存在
// alibaba.aliqin.fc.iot.device.isexist
//
// 判断设备是否存在
func AlibabaAliqinFcIotDeviceIsexist(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotDeviceIsexistAPIRequest, resp *aliqin.AlibabaAliqinFcIotDeviceIsexistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
