package alscmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// Alibabaalscmerchantextticketcodeuse 外部核销服务
// alibaba.alsc.merchant.ext.ticketcode.use
//
// 外部核销服务
func Alibabaalscmerchantextticketcodeuse(clt *core.SDKClient, req *alscmerchant.AlibabaalscmerchantextticketcodeuseAPIRequest, session string) (*alscmerchant.AlibabaalscmerchantextticketcodeuseAPIResponse, error) {
	var resp alscmerchant.AlibabaalscmerchantextticketcodeuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
