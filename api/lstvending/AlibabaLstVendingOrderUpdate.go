package lstvending

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// AlibabaLstVendingOrderUpdate 自动售货机订单物流信息回传
// alibaba.lst.vending.order.update
//
// 零售通与设备供应商进行订单对接，通过此接口回流订单物流信息。
func AlibabaLstVendingOrderUpdate(ctx context.Context, clt *core.SDKClient, req *lstvending.AlibabaLstVendingOrderUpdateAPIRequest, resp *lstvending.AlibabaLstVendingOrderUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
