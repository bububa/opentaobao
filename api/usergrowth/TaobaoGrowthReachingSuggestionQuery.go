package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingSuggestionQuery 厂商生态推荐信息查询
// taobao.growth.reaching.suggestion.query
//
// 厂商生态推荐信息查询
func TaobaoGrowthReachingSuggestionQuery(clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingSuggestionQueryAPIRequest, session string) (*usergrowth.TaobaoGrowthReachingSuggestionQueryAPIResponse, error) {
	var resp usergrowth.TaobaoGrowthReachingSuggestionQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
