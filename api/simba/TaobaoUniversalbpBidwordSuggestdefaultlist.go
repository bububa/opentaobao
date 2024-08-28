package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpBidwordSuggestdefaultlist 建议默认关键词
// taobao.universalbp.bidword.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词
func TaobaoUniversalbpBidwordSuggestdefaultlist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest, resp *simba.TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
