package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeaseExceptionflowsynchronize 天猫开新车租后异常流线下处理状态通知接口
// tmall.car.lease.exceptionflowsynchronize
//
// 天猫开新车租后异常流线下处理状态通知接口
func TmallCarLeaseExceptionflowsynchronize(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarLeaseExceptionflowsynchronizeAPIRequest, resp *tmallcar.TmallCarLeaseExceptionflowsynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
