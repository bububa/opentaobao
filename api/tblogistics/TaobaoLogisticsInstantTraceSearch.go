package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsInstantTraceSearch 物流详情查询
// taobao.logistics.instant.trace.search
//
// 物流详情查询
func TaobaoLogisticsInstantTraceSearch(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsInstantTraceSearchAPIRequest, resp *tblogistics.TaobaoLogisticsInstantTraceSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
