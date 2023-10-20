package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelTradeTemplateQuery 订单服务详情模版查询
// alitrip.travel.trade.template.query
//
// 通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息
func AlitripTravelTradeTemplateQuery(clt *core.SDKClient, req *traveltrade.AlitripTravelTradeTemplateQueryAPIRequest, resp *traveltrade.AlitripTravelTradeTemplateQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
