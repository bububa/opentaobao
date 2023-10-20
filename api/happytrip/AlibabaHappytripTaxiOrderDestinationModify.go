package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderDestinationModify 修改目的地
// alibaba.happytrip.taxi.order.destination.modify
//
// 通知ISV修改订单信息
func AlibabaHappytripTaxiOrderDestinationModify(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderDestinationModifyAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderDestinationModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
