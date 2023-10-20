package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketresend 外部合作商家重发电子凭证回调接口
// taobao.vmarket.eticket.resend
//
// 外部合作商家重发电子凭证回调接口
func Taobaovmarketeticketresend(clt *core.SDKClient, req *eticket.TaobaovmarketeticketresendAPIRequest, session string) (*eticket.TaobaovmarketeticketresendAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketresendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
