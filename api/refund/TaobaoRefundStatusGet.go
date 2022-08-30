package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRefundStatusGet 查询退款状态
// taobao.refund.status.get
//
// 根据主订单或者子订单id查询退款状态,入参传入主订单或者子订单号1.如果当传入子订单时，返回子订单最后一笔退款单的状态,如果子订单申请退款退款返回空list.2.如果传传入主订单，则返回主订单下所有子订单的最后一笔退款单状态，如果对应的子订单没有生成退款单，则对应子订单对应数据返回。
func TaobaoRefundStatusGet(clt *core.SDKClient, req *refund.TaobaoRefundStatusGetAPIRequest, session string) (*refund.TaobaoRefundStatusGetAPIResponse, error) {
	var resp refund.TaobaoRefundStatusGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
