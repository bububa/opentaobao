package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketTimeExpand 订单延时接口
// taobao.vmarket.eticket.time.expand
//
// 提供码商操作订单延期接口
func TaobaoVmarketEticketTimeExpand(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketTimeExpandAPIRequest, resp *eticket.TaobaoVmarketEticketTimeExpandAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
