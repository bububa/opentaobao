package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangCombinescitemBatchUpdate 更新组合货品
// alibaba.dchain.aoxiang.combinescitem.batch.update
//
// 更新组合货品
func AlibabaDchainAoxiangCombinescitemBatchUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangCombinescitemBatchUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
