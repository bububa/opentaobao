package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthReserveDentalMarkitem 标记商品是否可预约
// alibaba.alihealth.reserve.dental.markitem
//
// 标记商品是否可预约
func AlibabaAlihealthReserveDentalMarkitem(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalMarkitemAPIRequest, resp *alihealth2.AlibabaAlihealthReserveDentalMarkitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
