package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiUpdate 电子面单云打印更新接口
// cainiao.waybill.ii.update
//
// 商家更新电子面单号对应的面单信息。
func CainiaoWaybillIiUpdate(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillIiUpdateAPIRequest, resp *waybill.CainiaoWaybillIiUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
