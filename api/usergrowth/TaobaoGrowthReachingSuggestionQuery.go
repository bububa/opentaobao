package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingSuggestionQuery 厂商生态推荐信息查询
// taobao.growth.reaching.suggestion.query
//
// 厂商生态推荐信息查询
func TaobaoGrowthReachingSuggestionQuery(clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingSuggestionQueryAPIRequest, resp *usergrowth.TaobaoGrowthReachingSuggestionQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
