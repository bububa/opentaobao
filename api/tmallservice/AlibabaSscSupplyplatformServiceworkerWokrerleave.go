package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerWokrerleave 工人请假
// alibaba.ssc.supplyplatform.serviceworker.wokrerleave
//
// 工人请假
func AlibabaSscSupplyplatformServiceworkerWokrerleave(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
