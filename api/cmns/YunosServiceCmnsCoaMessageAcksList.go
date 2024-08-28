package cmns

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaMessageAcksList 消息ack记录查询
// yunos.service.cmns.coa.message.acks.list
//
// 第三方应用开发者调用此接口查询消息ack记录
func YunosServiceCmnsCoaMessageAcksList(ctx context.Context, clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageAcksListAPIRequest, resp *cmns.YunosServiceCmnsCoaMessageAcksListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
