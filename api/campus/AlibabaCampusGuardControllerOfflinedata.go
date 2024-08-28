package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardControllerOfflinedata 点位离线数据拉取
// alibaba.campus.guard.controller.offlinedata
//
// 点位离线数据拉取
func AlibabaCampusGuardControllerOfflinedata(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusGuardControllerOfflinedataAPIRequest, resp *campus.AlibabaCampusGuardControllerOfflinedataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
