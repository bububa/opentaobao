package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerCancelleave 工人取消请假
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
func AlibabaSscSupplyplatformServiceworkerCancelleave(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
