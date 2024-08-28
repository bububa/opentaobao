package cmns

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaMessageCancel CMNS消息撤回
// yunos.service.cmns.coa.message.cancel
//
// 此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。
func YunosServiceCmnsCoaMessageCancel(ctx context.Context, clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageCancelAPIRequest, resp *cmns.YunosServiceCmnsCoaMessageCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
