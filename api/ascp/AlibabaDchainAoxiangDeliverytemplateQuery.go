package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangDeliverytemplateQuery 商家运费模板查询
// alibaba.dchain.aoxiang.deliverytemplate.query
//
// 商家运费模板查询
func AlibabaDchainAoxiangDeliverytemplateQuery(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest, resp *ascp.AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
