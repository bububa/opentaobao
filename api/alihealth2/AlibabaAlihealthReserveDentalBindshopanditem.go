package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthReserveDentalBindshopanditem 绑定门店信息，商品信息
// alibaba.alihealth.reserve.dental.bindshopanditem
//
// 绑定门店信息，商品信息
func AlibabaAlihealthReserveDentalBindshopanditem(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalBindshopanditemAPIRequest, resp *alihealth2.AlibabaAlihealthReserveDentalBindshopanditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
