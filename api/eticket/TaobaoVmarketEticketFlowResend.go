package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketflowresend 业务重新触发发码短信
// taobao.vmarket.eticket.flow.resend
//
// 业务重新触发发码短信
func Taobaovmarketeticketflowresend(clt *core.SDKClient, req *eticket.TaobaovmarketeticketflowresendAPIRequest, session string) (*eticket.TaobaovmarketeticketflowresendAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketflowresendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
