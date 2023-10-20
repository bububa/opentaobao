package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaInsightCatsforecastnewGet 获取词的相关类目预测数据
// taobao.simba.insight.catsforecastnew.get
//
// 根据给定的词，预测这些词的相关类目
func TaobaoSimbaInsightCatsforecastnewGet(clt *core.SDKClient, req *simba.TaobaoSimbaInsightCatsforecastnewGetAPIRequest, resp *simba.TaobaoSimbaInsightCatsforecastnewGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
