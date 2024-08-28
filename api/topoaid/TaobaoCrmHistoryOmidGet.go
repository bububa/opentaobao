package topoaid

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoCrmHistoryOmidGet 根据buyerNick获取omid
// taobao.crm.history.omid.get
//
// 根据buyerNick获取ouid
func TaobaoCrmHistoryOmidGet(ctx context.Context, clt *core.SDKClient, req *topoaid.TaobaoCrmHistoryOmidGetAPIRequest, resp *topoaid.TaobaoCrmHistoryOmidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
