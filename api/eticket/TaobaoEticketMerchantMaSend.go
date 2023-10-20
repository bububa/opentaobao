package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaSend 码商发码成功回调接口
// taobao.eticket.merchant.ma.send
//
// 码商发码成功回调接口
func TaobaoEticketMerchantMaSend(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaSendAPIRequest, resp *eticket.TaobaoEticketMerchantMaSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
