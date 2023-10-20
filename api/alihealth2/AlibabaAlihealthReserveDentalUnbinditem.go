package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthreservedentalunbinditem 解绑商品信息
// alibaba.alihealth.reserve.dental.unbinditem
//
// 绑定门店信息，商品信息
func Alibabaalihealthreservedentalunbinditem(clt *core.SDKClient, req *alihealth2.AlibabaalihealthreservedentalunbinditemAPIRequest, session string) (*alihealth2.AlibabaalihealthreservedentalunbinditemAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthreservedentalunbinditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
