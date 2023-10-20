package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketflowconsume 无交易类凭证核销
// taobao.vmarket.eticket.flow.consume
//
// 无交易类凭证核销
func Taobaovmarketeticketflowconsume(clt *core.SDKClient, req *eticket.TaobaovmarketeticketflowconsumeAPIRequest, session string) (*eticket.TaobaovmarketeticketflowconsumeAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketflowconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
