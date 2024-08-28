package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverContentQuery 查询大包详情
// cainiao.global.handover.content.query
//
// 查询大包详情
func CainiaoGlobalHandoverContentQuery(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverContentQueryAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverContentQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
