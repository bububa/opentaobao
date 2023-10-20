package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgPunishOrderGet 淘宝客-推广者-处罚订单查询
// taobao.tbk.dg.punish.order.get
//
// 新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的
func TaobaoTbkDgPunishOrderGet(clt *core.SDKClient, req *tbk.TaobaoTbkDgPunishOrderGetAPIRequest, resp *tbk.TaobaoTbkDgPunishOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
