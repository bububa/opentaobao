package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkCouponContractCreate 营销券合同创建接口
// alibaba.wdk.coupon.contract.create
//
// 营销券合同创建接口
func AlibabaWdkCouponContractCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkCouponContractCreateAPIRequest, resp *wdk.AlibabaWdkCouponContractCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
