package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiorderconfirm 费用确认
// alibaba.happytrip.taxi.order.confirm
//
// 1.司机点结束计费,欢行会收到正常支付待评论 回调,确认费用无误欢行可以通过此接口确认并支付。
// 2.如果欢行一直不调用此接口,订单会在48小时后自动支付。
func Alibabahappytriptaxiorderconfirm(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiorderconfirmAPIRequest, session string) (*happytrip.AlibabahappytriptaxiorderconfirmAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
