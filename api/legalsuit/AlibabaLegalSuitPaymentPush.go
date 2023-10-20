package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitPaymentPush 外部推送缴费
// alibaba.legal.suit.payment.push
//
// 外部推送缴费
func AlibabaLegalSuitPaymentPush(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitPaymentPushAPIRequest, session string) (*legalsuit.AlibabaLegalSuitPaymentPushAPIResponse, error) {
	var resp legalsuit.AlibabaLegalSuitPaymentPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
