package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceOsupgradedetail 获取系统升级详情
// yunos.tvpubadmin.device.osupgradedetail
//
// 获取系统升级详情
func YunosTvpubadminDeviceOsupgradedetail(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceOsupgradedetailAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceOsupgradedetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
