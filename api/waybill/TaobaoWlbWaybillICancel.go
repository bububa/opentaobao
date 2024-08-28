package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// TaobaoWlbWaybillICancel 商家取消获取的电子面单号v1.0
// taobao.wlb.waybill.i.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
func TaobaoWlbWaybillICancel(ctx context.Context, clt *core.SDKClient, req *waybill.TaobaoWlbWaybillICancelAPIRequest, resp *waybill.TaobaoWlbWaybillICancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
