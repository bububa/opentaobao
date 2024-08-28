package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkPosAfterbuyBenefitSend 生态pos购后发放权益
// alibaba.wdk.pos.afterbuy.benefit.send
//
// 生态pos购后发放权益接口开放
func AlibabaWdkPosAfterbuyBenefitSend(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkPosAfterbuyBenefitSendAPIRequest, resp *wdk.AlibabaWdkPosAfterbuyBenefitSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
