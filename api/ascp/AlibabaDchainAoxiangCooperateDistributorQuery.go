package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangCooperateDistributorQuery 商家关系查询分销商
// alibaba.dchain.aoxiang.cooperate.distributor.query
//
// 商家关系查询分销商
func AlibabaDchainAoxiangCooperateDistributorQuery(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest, resp *ascp.AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
