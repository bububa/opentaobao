package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsTraceSearch 物流流转信息查询
// taobao.logistics.trace.search
//
// 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。&lt;br/&gt;此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为&lt;status&gt;对方已签收&lt;/status&gt;，该字段仅对线上发货（online.send）的运单有效。）
func TaobaoLogisticsTraceSearch(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsTraceSearchAPIRequest, resp *tblogistics.TaobaoLogisticsTraceSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
