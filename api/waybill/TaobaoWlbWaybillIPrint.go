package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// TaobaoWlbWaybillIPrint 打印确认接口v1.0
// taobao.wlb.waybill.i.print
//
// 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
func TaobaoWlbWaybillIPrint(ctx context.Context, clt *core.SDKClient, req *waybill.TaobaoWlbWaybillIPrintAPIRequest, resp *waybill.TaobaoWlbWaybillIPrintAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
