package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaInsightWordspricedataGet
关键词按竞价区间的分布数据
taobao.simba.insight.wordspricedata.get

获取关键词按竞价区间进行细分的数据 */
func TaobaoSimbaInsightWordspricedataGet(clt *core.SDKClient, req *simba.TaobaoSimbaInsightWordspricedataGetAPIRequest, session string) (*simba.TaobaoSimbaInsightWordspricedataGetAPIResponse, error) {
	var resp simba.TaobaoSimbaInsightWordspricedataGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
