package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardActive 标准激活卡
// alibaba.alsc.crm.card.active
//
// 激活卡
func AlibabaAlscCrmCardActive(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardActiveAPIRequest, resp *alsc.AlibabaAlscCrmCardActiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
