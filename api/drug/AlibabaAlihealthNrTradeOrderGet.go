package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrTradeOrderGet 获取订单详情
// alibaba.alihealth.nr.trade.order.get
//
// 阿里健康O2O，获取订单详情
func AlibabaAlihealthNrTradeOrderGet(clt *core.SDKClient, req *drug.AlibabaAlihealthNrTradeOrderGetAPIRequest, resp *drug.AlibabaAlihealthNrTradeOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
