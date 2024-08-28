package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// AlibabaRetailCommissionResultQuery 分佣结果查询
// alibaba.retail.commission.result.query
//
// 查询导购分佣记录
func AlibabaRetailCommissionResultQuery(ctx context.Context, clt *core.SDKClient, req *omniorder.AlibabaRetailCommissionResultQueryAPIRequest, resp *omniorder.AlibabaRetailCommissionResultQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
