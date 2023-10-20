package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// Alibabaalscdaodianticketconsult 券码预览
// alibaba.alsc.daodian.ticket.consult
//
// 券码预览
func Alibabaalscdaodianticketconsult(clt *core.SDKClient, req *alscmerchant.AlibabaalscdaodianticketconsultAPIRequest, session string) (*alscmerchant.AlibabaalscdaodianticketconsultAPIResponse, error) {
	var resp alscmerchant.AlibabaalscdaodianticketconsultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
