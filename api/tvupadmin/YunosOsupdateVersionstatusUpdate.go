package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateVersionstatusUpdate 更新应用升级状态
// yunos.osupdate.versionstatus.update
//
// 更新应用升级状态
func YunosOsupdateVersionstatusUpdate(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosOsupdateVersionstatusUpdateAPIRequest, resp *tvupadmin.YunosOsupdateVersionstatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
