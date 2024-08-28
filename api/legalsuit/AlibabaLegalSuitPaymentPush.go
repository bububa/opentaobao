package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitPaymentPush 外部推送缴费
// alibaba.legal.suit.payment.push
//
// 外部推送缴费
func AlibabaLegalSuitPaymentPush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitPaymentPushAPIRequest, resp *legalsuit.AlibabaLegalSuitPaymentPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
