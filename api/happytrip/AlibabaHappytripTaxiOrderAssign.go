package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderAssign 订单指派
// alibaba.happytrip.taxi.order.assign
//
// 通知供应商订单指派成功
func AlibabaHappytripTaxiOrderAssign(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderAssignAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderAssignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
