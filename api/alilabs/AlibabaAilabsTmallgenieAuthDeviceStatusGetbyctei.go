package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGetbyctei 根据ctei查询状态
// alibaba.ailabs.tmallgenie.auth.device.status.getbyctei
//
// 提供给电信查询设备在线状态值
func AlibabaAilabsTmallgenieAuthDeviceStatusGetbyctei(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
