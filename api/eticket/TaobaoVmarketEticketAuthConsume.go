package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketauthconsume 核销放行的核销接口
// taobao.vmarket.eticket.auth.consume
//
// 针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口
func Taobaovmarketeticketauthconsume(clt *core.SDKClient, req *eticket.TaobaovmarketeticketauthconsumeAPIRequest, session string) (*eticket.TaobaovmarketeticketauthconsumeAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketauthconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
