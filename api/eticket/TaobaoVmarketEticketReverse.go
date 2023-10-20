package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketreverse 电子凭证冲正接口
// taobao.vmarket.eticket.reverse
//
// 电子凭证平台冲正接口
func Taobaovmarketeticketreverse(clt *core.SDKClient, req *eticket.TaobaovmarketeticketreverseAPIRequest, session string) (*eticket.TaobaovmarketeticketreverseAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketreverseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
