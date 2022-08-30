package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscMerchantExtTicketcodeSend 异步发码
// alibaba.alsc.merchant.ext.ticketcode.send
//
// 外部券异步发码
func AlibabaAlscMerchantExtTicketcodeSend(clt *core.SDKClient, req *alscmerchant.AlibabaAlscMerchantExtTicketcodeSendAPIRequest, session string) (*alscmerchant.AlibabaAlscMerchantExtTicketcodeSendAPIResponse, error) {
	var resp alscmerchant.AlibabaAlscMerchantExtTicketcodeSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
