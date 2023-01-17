package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscDaodianTicketConsult 券码预览
// alibaba.alsc.daodian.ticket.consult
//
// 券码预览
func AlibabaAlscDaodianTicketConsult(clt *core.SDKClient, req *alscmerchant.AlibabaAlscDaodianTicketConsultAPIRequest, session string) (*alscmerchant.AlibabaAlscDaodianTicketConsultAPIResponse, error) {
	var resp alscmerchant.AlibabaAlscDaodianTicketConsultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
