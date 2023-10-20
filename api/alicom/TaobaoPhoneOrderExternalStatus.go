package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// TaobaoPhoneOrderExternalStatus 话费外放订单状态接口
// taobao.phone.order.external.status
//
// 话费外放订单状态接口
func TaobaoPhoneOrderExternalStatus(clt *core.SDKClient, req *alicom.TaobaoPhoneOrderExternalStatusAPIRequest, session string) (*alicom.TaobaoPhoneOrderExternalStatusAPIResponse, error) {
	var resp alicom.TaobaoPhoneOrderExternalStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
