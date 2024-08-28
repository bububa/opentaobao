package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemmappingBatchCreate 新建商货品关联
// alibaba.dchain.aoxiang.itemmapping.batch.create
//
// 新建商货品关联
func AlibabaDchainAoxiangItemmappingBatchCreate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemmappingBatchCreateAPIRequest, resp *ascp.AlibabaDchainAoxiangItemmappingBatchCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
