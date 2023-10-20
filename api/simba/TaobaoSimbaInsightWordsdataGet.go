package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaInsightWordsdataGet 获取关键词的大盘数据
// taobao.simba.insight.wordsdata.get
//
// 获取关键词的详细数据
func TaobaoSimbaInsightWordsdataGet(clt *core.SDKClient, req *simba.TaobaoSimbaInsightWordsdataGetAPIRequest, resp *simba.TaobaoSimbaInsightWordsdataGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
