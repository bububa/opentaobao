package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrTradeOrderGet 获取订单详情
// alibaba.alihealth.nr.trade.order.get
//
// 阿里健康O2O，获取订单详情
func AlibabaAlihealthNrTradeOrderGet(clt *core.SDKClient, req *drug.AlibabaAlihealthNrTradeOrderGetAPIRequest, session string) (*drug.AlibabaAlihealthNrTradeOrderGetAPIResponse, error) {
	var resp drug.AlibabaAlihealthNrTradeOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
