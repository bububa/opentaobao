package usergrowth

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingSuggestionQuery 厂商生态推荐信息查询
// taobao.growth.reaching.suggestion.query
//
// 厂商生态推荐信息查询
func TaobaoGrowthReachingSuggestionQuery(ctx context.Context, clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingSuggestionQueryAPIRequest, resp *usergrowth.TaobaoGrowthReachingSuggestionQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
