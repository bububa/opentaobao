package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderGet 订单详情
// alibaba.happytrip.taxi.order.get
//
// 获取订单状态及详情
func AlibabaHappytripTaxiOrderGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderGetAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
