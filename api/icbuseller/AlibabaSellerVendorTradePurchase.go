package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// Alibabasellervendortradepurchase 查看购买人的订单记录以及授权时间
// alibaba.seller.vendor.trade.purchase
//
// 查看购买人的订单记录以及授权时间
func Alibabasellervendortradepurchase(clt *core.SDKClient, req *icbuseller.AlibabasellervendortradepurchaseAPIRequest, session string) (*icbuseller.AlibabasellervendortradepurchaseAPIResponse, error) {
	var resp icbuseller.AlibabasellervendortradepurchaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
