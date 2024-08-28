package maitix

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixOrderConfirm 大麦-出票
// alibaba.damai.maitix.order.confirm
//
// 出票
func AlibabaDamaiMaitixOrderConfirm(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderConfirmAPIRequest, resp *maitix.AlibabaDamaiMaitixOrderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
