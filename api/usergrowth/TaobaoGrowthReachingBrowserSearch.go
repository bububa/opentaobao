package usergrowth

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingBrowserSearch 查询搜索关联
// taobao.growth.reaching.browser.search
//
// 查询搜索关联
func TaobaoGrowthReachingBrowserSearch(ctx context.Context, clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingBrowserSearchAPIRequest, resp *usergrowth.TaobaoGrowthReachingBrowserSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
