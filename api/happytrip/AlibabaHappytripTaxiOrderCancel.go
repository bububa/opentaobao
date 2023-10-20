package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderCancel 取消叫车
// alibaba.happytrip.taxi.order.cancel
//
// 取消叫车订单,行程中的订单不能取消
func AlibabaHappytripTaxiOrderCancel(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderCancelAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
