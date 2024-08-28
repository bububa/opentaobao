package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaTianjiDistributorOrderSubmit 分销商提交受理订单
// alibaba.tianji.distributor.order.submit
//
// 分销商提交受理订单，如合约订购、充值受理等
func AlibabaTianjiDistributorOrderSubmit(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaTianjiDistributorOrderSubmitAPIRequest, resp *alicom.AlibabaTianjiDistributorOrderSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
