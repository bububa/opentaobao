package usergrowth

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingXiniaoQuery 查询溪鸟推荐信息数据
// taobao.growth.reaching.xiniao.query
//
// 查询溪鸟推荐信息数据
func TaobaoGrowthReachingXiniaoQuery(ctx context.Context, clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingXiniaoQueryAPIRequest, resp *usergrowth.TaobaoGrowthReachingXiniaoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
