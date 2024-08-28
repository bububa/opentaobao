package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceQuery 获取设备列表
// yunos.tvpubadmin.device.query
//
// 获取设备列表
func YunosTvpubadminDeviceQuery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceQueryAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
