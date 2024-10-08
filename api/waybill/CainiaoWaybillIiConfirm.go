package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiConfirm 物流订单确认接口
// cainiao.waybill.ii.confirm
//
// 物流订单确认
func CainiaoWaybillIiConfirm(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillIiConfirmAPIRequest, resp *waybill.CainiaoWaybillIiConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
