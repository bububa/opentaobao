package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// TaobaoPhoneOrderExternalCreate 数字虚拟话费外放下单接口
// taobao.phone.order.external.create
//
// 数字虚拟话费外放下单接口
func TaobaoPhoneOrderExternalCreate(clt *core.SDKClient, req *alicom.TaobaoPhoneOrderExternalCreateAPIRequest, session string) (*alicom.TaobaoPhoneOrderExternalCreateAPIResponse, error) {
	var resp alicom.TaobaoPhoneOrderExternalCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
