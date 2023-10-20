package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// Alibabaalscmerchantextticketcodesend 异步发码
// alibaba.alsc.merchant.ext.ticketcode.send
//
// 外部券异步发码
func Alibabaalscmerchantextticketcodesend(clt *core.SDKClient, req *alscmerchant.AlibabaalscmerchantextticketcodesendAPIRequest, session string) (*alscmerchant.AlibabaalscmerchantextticketcodesendAPIResponse, error) {
	var resp alscmerchant.AlibabaalscmerchantextticketcodesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
