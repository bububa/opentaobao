package elife

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/elife"
)

// TaobaoElifeLifecardQuery 查询交易结果
// taobao.elife.lifecard.query
//
// 卖家在交易状态不明的情况下, 查询交易结果.
func TaobaoElifeLifecardQuery(ctx context.Context, clt *core.SDKClient, req *elife.TaobaoElifeLifecardQueryAPIRequest, resp *elife.TaobaoElifeLifecardQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
