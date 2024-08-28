package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaInsightCatsinfoGet 类目信息获取
// taobao.simba.insight.catsinfo.get
//
// 获取类目信息，此接口既提供所有顶级类目的查询，又提供给定类目id自身信息和子类目信息的查询，所以可以根据此接口逐层获取所有的类目信息
func TaobaoSimbaInsightCatsinfoGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaInsightCatsinfoGetAPIRequest, resp *simba.TaobaoSimbaInsightCatsinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
