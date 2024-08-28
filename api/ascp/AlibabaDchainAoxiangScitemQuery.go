package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangScitemQuery 货品查询
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
func AlibabaDchainAoxiangScitemQuery(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangScitemQueryAPIRequest, resp *ascp.AlibabaDchainAoxiangScitemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
