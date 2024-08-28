package cmns

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaMessageGet 消息详情查询
// yunos.service.cmns.coa.message.get
//
// 第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息
func YunosServiceCmnsCoaMessageGet(ctx context.Context, clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageGetAPIRequest, resp *cmns.YunosServiceCmnsCoaMessageGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
