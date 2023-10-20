package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusRefundticketpriceSet 汽车票退款申请接口
// taobao.bus.refundticketprice.set
//
// 汽车票代理商利用该接口申请退票
func TaobaoBusRefundticketpriceSet(clt *core.SDKClient, req *bus.TaobaoBusRefundticketpriceSetAPIRequest, session string) (*bus.TaobaoBusRefundticketpriceSetAPIResponse, error) {
	var resp bus.TaobaoBusRefundticketpriceSetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
