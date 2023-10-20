package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitPaymentPush 外部推送缴费
// alibaba.legal.suit.payment.push
//
// 外部推送缴费
func AlibabaLegalSuitPaymentPush(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitPaymentPushAPIRequest, resp *legalsuit.AlibabaLegalSuitPaymentPushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
