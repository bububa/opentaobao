package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderNotify 状态通知
// alibaba.happytrip.taxi.order.notify
//
// 当订单发生变化是供应商通过状态通知API通知欢行，欢行获取最新的订单详情和状态进行业务处理。
func AlibabaHappytripTaxiOrderNotify(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderNotifyAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
