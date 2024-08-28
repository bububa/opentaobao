package cmns

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaMessageresultGet CMNS消息发送到达查询
// yunos.service.cmns.coa.messageresult.get
//
// CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
func YunosServiceCmnsCoaMessageresultGet(ctx context.Context, clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageresultGetAPIRequest, resp *cmns.YunosServiceCmnsCoaMessageresultGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
