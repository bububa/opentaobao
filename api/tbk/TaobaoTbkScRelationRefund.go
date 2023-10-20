package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScRelationRefund 淘宝客-服务商-维权退款订单查询
// taobao.tbk.sc.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
func TaobaoTbkScRelationRefund(clt *core.SDKClient, req *tbk.TaobaoTbkScRelationRefundAPIRequest, resp *tbk.TaobaoTbkScRelationRefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
