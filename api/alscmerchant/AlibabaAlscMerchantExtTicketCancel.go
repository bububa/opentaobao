package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscMerchantExtTicketCancel 凭证作废
// alibaba.alsc.merchant.ext.ticket.cancel
//
// isv调用银行接口超时导致凭证信息同步失败,触发关单处理,调用作废接口
func AlibabaAlscMerchantExtTicketCancel(clt *core.SDKClient, req *alscmerchant.AlibabaAlscMerchantExtTicketCancelAPIRequest, session string) (*alscmerchant.AlibabaAlscMerchantExtTicketCancelAPIResponse, error) {
	var resp alscmerchant.AlibabaAlscMerchantExtTicketCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
