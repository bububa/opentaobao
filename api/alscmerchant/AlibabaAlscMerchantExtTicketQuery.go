package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscMerchantExtTicketQuery 口碑凭证模版查询服务
// alibaba.alsc.merchant.ext.ticket.query
//
// isv需要在c端展示凭证信息,需要查询凭证模版
func AlibabaAlscMerchantExtTicketQuery(clt *core.SDKClient, req *alscmerchant.AlibabaAlscMerchantExtTicketQueryAPIRequest, session string) (*alscmerchant.AlibabaAlscMerchantExtTicketQueryAPIResponse, error) {
	var resp alscmerchant.AlibabaAlscMerchantExtTicketQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
