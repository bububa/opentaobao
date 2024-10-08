package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaXianyuTenderOrderPerform 闲鱼暗拍订单履约
// alibaba.xianyu.tender.order.perform
//
// 闲鱼暗拍订单履约
func AlibabaXianyuTenderOrderPerform(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaXianyuTenderOrderPerformAPIRequest, resp *idle.AlibabaXianyuTenderOrderPerformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
