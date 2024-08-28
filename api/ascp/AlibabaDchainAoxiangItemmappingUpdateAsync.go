package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemmappingUpdateAsync 创建/更新商货品关联关系
// alibaba.dchain.aoxiang.itemmapping.update.async
//
// 创建/更新商货品关联关系
func AlibabaDchainAoxiangItemmappingUpdateAsync(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest, resp *ascp.AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
