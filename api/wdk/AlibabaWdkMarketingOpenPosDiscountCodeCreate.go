package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingOpenPosDiscountCodeCreate pos一物一码创建
// alibaba.wdk.marketing.open.pos.discount.code.create
//
// pos一物一码创建
func AlibabaWdkMarketingOpenPosDiscountCodeCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest, resp *wdk.AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
