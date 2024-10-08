package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoUopIntercept 阿里巴巴.天猫. 履约订单. 配送拦截
// alibaba.tianmao.uop.intercept
//
// 阿里巴巴.天猫. 履约订单. 配送拦截
func AlibabaTianmaoUopIntercept(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaTianmaoUopInterceptAPIRequest, resp *ascp.AlibabaTianmaoUopInterceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
