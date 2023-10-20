package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthreservedentalbindshopanditem 绑定门店信息，商品信息
// alibaba.alihealth.reserve.dental.bindshopanditem
//
// 绑定门店信息，商品信息
func Alibabaalihealthreservedentalbindshopanditem(clt *core.SDKClient, req *alihealth2.AlibabaalihealthreservedentalbindshopanditemAPIRequest, session string) (*alihealth2.AlibabaalihealthreservedentalbindshopanditemAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthreservedentalbindshopanditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
