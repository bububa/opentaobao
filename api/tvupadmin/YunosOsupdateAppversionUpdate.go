package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateAppversionUpdate 应用升级任务更新
// yunos.osupdate.appversion.update
//
// 应用升级任务更新
func YunosOsupdateAppversionUpdate(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionUpdateAPIRequest, resp *tvupadmin.YunosOsupdateAppversionUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
