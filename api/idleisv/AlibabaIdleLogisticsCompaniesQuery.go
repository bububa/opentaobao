package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleLogisticsCompaniesQuery 快递公司列表查询
// alibaba.idle.logistics.companies.query
//
// 支持发货的快递公司列表查询
func AlibabaIdleLogisticsCompaniesQuery(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleLogisticsCompaniesQueryAPIRequest, resp *idleisv.AlibabaIdleLogisticsCompaniesQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
