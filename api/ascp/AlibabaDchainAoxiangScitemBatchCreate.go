package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemBatchCreate 新建货品
// alibaba.dchain.aoxiang.scitem.batch.create
//
// 新建货品
func AlibabaDchainAoxiangScitemBatchCreate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemBatchCreateAPIRequest, resp *ascp.AlibabaDchainAoxiangScitemBatchCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
