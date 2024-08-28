package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemDelete 货品删除
// alibaba.dchain.aoxiang.scitem.delete
//
// 货品删除
func AlibabaDchainAoxiangScitemDelete(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemDeleteAPIRequest, resp *ascp.AlibabaDchainAoxiangScitemDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
