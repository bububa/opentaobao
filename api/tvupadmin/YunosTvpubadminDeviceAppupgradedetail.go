package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceAppupgradedetail 获取应用升级详情
// yunos.tvpubadmin.device.appupgradedetail
//
// 获取应用升级详情
func YunosTvpubadminDeviceAppupgradedetail(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceAppupgradedetailAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceAppupgradedetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
