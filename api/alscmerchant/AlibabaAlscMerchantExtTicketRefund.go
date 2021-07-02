package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscMerchantExtTicketRefund 口碑凭证售后退
// alibaba.alsc.merchant.ext.ticket.refund
//
// 口碑凭证售后退
func AlibabaAlscMerchantExtTicketRefund(clt *core.SDKClient, req *alscmerchant.AlibabaAlscMerchantExtTicketRefundAPIRequest, session string) (*alscmerchant.AlibabaAlscMerchantExtTicketRefundAPIResponse, error) {
	var resp alscmerchant.AlibabaAlscMerchantExtTicketRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
