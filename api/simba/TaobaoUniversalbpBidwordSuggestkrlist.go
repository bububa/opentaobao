package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpBidwordSuggestkrlist 关键词建议
// taobao.universalbp.bidword.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词
func TaobaoUniversalbpBidwordSuggestkrlist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpBidwordSuggestkrlistAPIRequest, resp *simba.TaobaoUniversalbpBidwordSuggestkrlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
