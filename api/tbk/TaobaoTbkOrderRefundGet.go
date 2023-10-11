package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkOrderRefundGet 淘宝客-推广者-全量维权退款订单查询
// taobao.tbk.order.refund.get
//
// 淘宝客账户下全量维权退款订单查询
func TaobaoTbkOrderRefundGet(clt *core.SDKClient, req *tbk.TaobaoTbkOrderRefundGetAPIRequest, session string) (*tbk.TaobaoTbkOrderRefundGetAPIResponse, error) {
	var resp tbk.TaobaoTbkOrderRefundGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
