package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsInstantTraceSearch 物流详情查询
// taobao.logistics.instant.trace.search
//
// 物流详情查询
func TaobaoLogisticsInstantTraceSearch(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsInstantTraceSearchAPIRequest, session string) (*tblogistics.TaobaoLogisticsInstantTraceSearchAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsInstantTraceSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
