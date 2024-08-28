package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaInsightCatstopwordnewGet 获取类目下最热门的词
// taobao.simba.insight.catstopwordnew.get
//
// 按照某个维度，查询某个类目下最热门的词，维度有点击，展现，花费，点击率等，具体可以按哪些字段进行排序，参考参数说明，比如选择了impression，则返回该类目下展现量最高那几个词。
func TaobaoSimbaInsightCatstopwordnewGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaInsightCatstopwordnewGetAPIRequest, resp *simba.TaobaoSimbaInsightCatstopwordnewGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
