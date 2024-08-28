package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceUpdateosstatus 更新系统版本审核状态
// yunos.tvpubadmin.device.updateosstatus
//
// 更新系统版本审核状态
func YunosTvpubadminDeviceUpdateosstatus(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceUpdateosstatusAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceUpdateosstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
