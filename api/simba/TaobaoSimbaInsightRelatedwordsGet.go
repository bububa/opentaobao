package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaInsightRelatedwordsGet 获取词的相关词
// taobao.simba.insight.relatedwords.get
//
// 获取给定词的若干相关词，返回结果中越相关的权重越大，排在越前面，根据number参数对返回结果进行截断。
func TaobaoSimbaInsightRelatedwordsGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaInsightRelatedwordsGetAPIRequest, resp *simba.TaobaoSimbaInsightRelatedwordsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
