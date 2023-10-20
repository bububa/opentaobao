package gameact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// TaobaoDeActivityDeliveryAddrConfirm 用户收件地址确认
// taobao.de.activity.delivery.addr.confirm
//
// 用户收件地址确认
func TaobaoDeActivityDeliveryAddrConfirm(clt *core.SDKClient, req *gameact.TaobaoDeActivityDeliveryAddrConfirmAPIRequest, resp *gameact.TaobaoDeActivityDeliveryAddrConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
