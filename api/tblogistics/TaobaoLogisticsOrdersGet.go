package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsOrdersGet 批量查询物流订单
// taobao.logistics.orders.get
//
// 批量查询物流订单。
func TaobaoLogisticsOrdersGet(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsOrdersGetAPIRequest, session string) (*tblogistics.TaobaoLogisticsOrdersGetAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsOrdersGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
