package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentDeviceGetvendor 查询设备Vendor信息
// yunos.tvpubadmin.content.device.getvendor
//
// 查询设备Vendor信息
func YunosTvpubadminContentDeviceGetvendor(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentDeviceGetvendorAPIRequest, resp *tvupadmin.YunosTvpubadminContentDeviceGetvendorAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
