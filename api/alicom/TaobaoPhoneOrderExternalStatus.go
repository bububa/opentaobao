package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// TaobaoPhoneOrderExternalStatus 话费外放订单状态接口
// taobao.phone.order.external.status
//
// 话费外放订单状态接口
func TaobaoPhoneOrderExternalStatus(clt *core.SDKClient, req *alicom.TaobaoPhoneOrderExternalStatusAPIRequest, resp *alicom.TaobaoPhoneOrderExternalStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
