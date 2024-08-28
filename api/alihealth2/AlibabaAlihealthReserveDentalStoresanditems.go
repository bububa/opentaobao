package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthReserveDentalStoresanditems 查询商户门店，商品列表
// alibaba.alihealth.reserve.dental.storesanditems
//
// 查询商户门店，商品列表
func AlibabaAlihealthReserveDentalStoresanditems(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalStoresanditemsAPIRequest, resp *alihealth2.AlibabaAlihealthReserveDentalStoresanditemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
