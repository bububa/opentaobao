package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscDaodianTicketReserve 外部券冲正
// alibaba.alsc.daodian.ticket.reserve
//
// 外部券冲正
func AlibabaAlscDaodianTicketReserve(clt *core.SDKClient, req *alscmerchant.AlibabaAlscDaodianTicketReserveAPIRequest, session string) (*alscmerchant.AlibabaAlscDaodianTicketReserveAPIResponse, error) {
	var resp alscmerchant.AlibabaAlscDaodianTicketReserveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
