package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateAppversionInfo 获取应用升级详情
// yunos.osupdate.appversion.info
//
// 获取应用升级详情
func YunosOsupdateAppversionInfo(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionInfoAPIRequest, resp *tvupadmin.YunosOsupdateAppversionInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
