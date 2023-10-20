package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaoeticketmerchantmafailsend 码商发码失败回调接口
// taobao.eticket.merchant.ma.failsend
//
// 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
func Taobaoeticketmerchantmafailsend(clt *core.SDKClient, req *eticket.TaobaoeticketmerchantmafailsendAPIRequest, session string) (*eticket.TaobaoeticketmerchantmafailsendAPIResponse, error) {
	var resp eticket.TaobaoeticketmerchantmafailsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
