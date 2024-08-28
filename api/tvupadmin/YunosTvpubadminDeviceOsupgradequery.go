package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceOsupgradequery 系统升级查询
// yunos.tvpubadmin.device.osupgradequery
//
// 系统升级查询
func YunosTvpubadminDeviceOsupgradequery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceOsupgradequeryAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceOsupgradequeryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
