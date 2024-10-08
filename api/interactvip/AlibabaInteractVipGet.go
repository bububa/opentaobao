package interactvip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interactvip"
)

// AlibabaInteractVipGet 会员淘气值获取
// alibaba.interact.vip.get
//
// 提供用户淘气值&amp;用户角色身份查询
func AlibabaInteractVipGet(ctx context.Context, clt *core.SDKClient, req *interactvip.AlibabaInteractVipGetAPIRequest, resp *interactvip.AlibabaInteractVipGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
