package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGet 设备状态查询
// alibaba.ailabs.tmallgenie.auth.device.status.get
//
// 提供给天猫精灵定制机厂商 查询设备在线状态值
func AlibabaAilabsTmallgenieAuthDeviceStatusGet(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
