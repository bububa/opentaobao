package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaoeticketmerchantmaresend 电子凭证重发回调接口
// taobao.eticket.merchant.ma.resend
//
// 码商重发电子凭证回调接口
func Taobaoeticketmerchantmaresend(clt *core.SDKClient, req *eticket.TaobaoeticketmerchantmaresendAPIRequest, session string) (*eticket.TaobaoeticketmerchantmaresendAPIResponse, error) {
	var resp eticket.TaobaoeticketmerchantmaresendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
