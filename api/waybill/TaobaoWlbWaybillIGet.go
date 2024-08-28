package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// TaobaoWlbWaybillIGet 获取物流服务商电子面单号v1.0
// taobao.wlb.waybill.i.get
//
// 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
func TaobaoWlbWaybillIGet(ctx context.Context, clt *core.SDKClient, req *waybill.TaobaoWlbWaybillIGetAPIRequest, resp *waybill.TaobaoWlbWaybillIGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
