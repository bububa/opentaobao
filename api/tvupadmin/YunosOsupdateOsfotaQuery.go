package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateOsfotaQuery 系统升级分页查询
// yunos.osupdate.osfota.query
//
// 分页查询osoupdate系统升级列表
func YunosOsupdateOsfotaQuery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosOsupdateOsfotaQueryAPIRequest, resp *tvupadmin.YunosOsupdateOsfotaQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
