package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemBatchUpdate alibaba.dchain.aoxiang.scitem.batch.update
// alibaba.dchain.aoxiang.scitem.batch.update
//
// 更新货品
func AlibabaDchainAoxiangScitemBatchUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemBatchUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangScitemBatchUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
