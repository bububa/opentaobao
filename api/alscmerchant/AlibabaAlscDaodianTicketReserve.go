package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// Alibabaalscdaodianticketreserve 外部券冲正
// alibaba.alsc.daodian.ticket.reserve
//
// 外部券冲正
func Alibabaalscdaodianticketreserve(clt *core.SDKClient, req *alscmerchant.AlibabaalscdaodianticketreserveAPIRequest, session string) (*alscmerchant.AlibabaalscdaodianticketreserveAPIResponse, error) {
	var resp alscmerchant.AlibabaalscdaodianticketreserveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
