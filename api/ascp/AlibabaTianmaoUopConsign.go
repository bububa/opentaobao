package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoUopConsign 阿里巴巴.天猫. 履约订单. 发货
// alibaba.tianmao.uop.consign
//
// 阿里巴巴.天猫. 履约订单. 发货
func AlibabaTianmaoUopConsign(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaTianmaoUopConsignAPIRequest, resp *ascp.AlibabaTianmaoUopConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
