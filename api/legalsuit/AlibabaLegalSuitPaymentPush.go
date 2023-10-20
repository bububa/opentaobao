package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitpaymentpush 外部推送缴费
// alibaba.legal.suit.payment.push
//
// 外部推送缴费
func Alibabalegalsuitpaymentpush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitpaymentpushAPIRequest, session string) (*legalsuit.AlibabalegalsuitpaymentpushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitpaymentpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
