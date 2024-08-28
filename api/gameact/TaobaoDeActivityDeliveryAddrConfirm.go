package gameact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// TaobaoDeActivityDeliveryAddrConfirm 用户收件地址确认
// taobao.de.activity.delivery.addr.confirm
//
// 用户收件地址确认
func TaobaoDeActivityDeliveryAddrConfirm(ctx context.Context, clt *core.SDKClient, req *gameact.TaobaoDeActivityDeliveryAddrConfirmAPIRequest, resp *gameact.TaobaoDeActivityDeliveryAddrConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
