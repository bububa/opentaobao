package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRefundsReceiveGet 查询卖家收到的退款列表
// taobao.refunds.receive.get
//
// 查询卖家收到的退款列表
func TaobaoRefundsReceiveGet(clt *core.SDKClient, req *refund.TaobaoRefundsReceiveGetAPIRequest, resp *refund.TaobaoRefundsReceiveGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
