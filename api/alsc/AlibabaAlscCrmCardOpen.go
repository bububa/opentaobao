package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardOpen 标准开卡流程
// alibaba.alsc.crm.card.open
//
// 标准开卡流程
func AlibabaAlscCrmCardOpen(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardOpenAPIRequest, resp *alsc.AlibabaAlscCrmCardOpenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
