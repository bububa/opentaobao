package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaInsightWordspricedataGet 关键词按竞价区间的分布数据
// taobao.simba.insight.wordspricedata.get
//
// 获取关键词按竞价区间进行细分的数据
func TaobaoSimbaInsightWordspricedataGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaInsightWordspricedataGetAPIRequest, resp *simba.TaobaoSimbaInsightWordspricedataGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
