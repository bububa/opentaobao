package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpBidwordFindlist 词列表查询
// taobao.universalbp.bidword.findlist
//
// 根据计划+单元id，查绑定的词列表
func TaobaoUniversalbpBidwordFindlist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpBidwordFindlistAPIRequest, resp *simba.TaobaoUniversalbpBidwordFindlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
