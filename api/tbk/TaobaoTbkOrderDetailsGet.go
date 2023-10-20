package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkOrderDetailsGet 淘宝客-推广者-所有订单查询
// taobao.tbk.order.details.get
//
// 淘宝客推广带来的所有拍下付款的正向订单明细报表。
func TaobaoTbkOrderDetailsGet(clt *core.SDKClient, req *tbk.TaobaoTbkOrderDetailsGetAPIRequest, resp *tbk.TaobaoTbkOrderDetailsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
