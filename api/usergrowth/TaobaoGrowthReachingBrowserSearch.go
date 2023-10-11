package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingBrowserSearch 查询搜索关联
// taobao.growth.reaching.browser.search
//
// 查询搜索关联
func TaobaoGrowthReachingBrowserSearch(clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingBrowserSearchAPIRequest, session string) (*usergrowth.TaobaoGrowthReachingBrowserSearchAPIResponse, error) {
	var resp usergrowth.TaobaoGrowthReachingBrowserSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
