package topoaid

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoCrmHistoryOuidGet 根据buyerNick获取ouid
// taobao.crm.history.ouid.get
//
// 根据buyerNick获取ouid
func TaobaoCrmHistoryOuidGet(ctx context.Context, clt *core.SDKClient, req *topoaid.TaobaoCrmHistoryOuidGetAPIRequest, resp *topoaid.TaobaoCrmHistoryOuidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
