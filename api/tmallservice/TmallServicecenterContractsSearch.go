package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterContractsSearch 获取合同类的服务工单信息
// tmall.servicecenter.contracts.search
//
// 获取合同类的服务工单信息
func TmallServicecenterContractsSearch(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterContractsSearchAPIRequest, resp *tmallservice.TmallServicecenterContractsSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
