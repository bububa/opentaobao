package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketauthbeforeconsume 核销放行的查询接口
// taobao.vmarket.eticket.auth.beforeconsume
//
// 针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销查询接口
func Taobaovmarketeticketauthbeforeconsume(clt *core.SDKClient, req *eticket.TaobaovmarketeticketauthbeforeconsumeAPIRequest, session string) (*eticket.TaobaovmarketeticketauthbeforeconsumeAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketauthbeforeconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
