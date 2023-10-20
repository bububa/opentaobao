package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantTbmaGet 码商查询淘宝码接口
// taobao.eticket.merchant.tbma.get
//
// 码商查询淘宝码接口
func TaobaoEticketMerchantTbmaGet(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantTbmaGetAPIRequest, resp *eticket.TaobaoEticketMerchantTbmaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
