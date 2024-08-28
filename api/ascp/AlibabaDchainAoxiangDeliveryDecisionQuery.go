package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangDeliveryDecisionQuery 查询黑白名单快递
// alibaba.dchain.aoxiang.delivery.decision.query
//
// 查询黑白名单快递
func AlibabaDchainAoxiangDeliveryDecisionQuery(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest, resp *ascp.AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
