package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthDeviceUnbind 解绑设备
// alibaba.ailabs.tmallgenie.auth.device.unbind
//
// 通过此接口解绑天猫精灵设备
func AlibabaAilabsTmallgenieAuthDeviceUnbind(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
