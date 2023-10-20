package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusrefundticketpriceset 汽车票退款申请接口
// taobao.bus.refundticketprice.set
//
// 汽车票代理商利用该接口申请退票
func Taobaobusrefundticketpriceset(clt *core.SDKClient, req *bus.TaobaobusrefundticketpricesetAPIRequest, session string) (*bus.TaobaobusrefundticketpricesetAPIResponse, error) {
	var resp bus.TaobaobusrefundticketpricesetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
