package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticsordersget 批量查询物流订单
// taobao.logistics.orders.get
//
// 批量查询物流订单。
func Taobaologisticsordersget(clt *core.SDKClient, req *tblogistics.TaobaologisticsordersgetAPIRequest, session string) (*tblogistics.TaobaologisticsordersgetAPIResponse, error) {
	var resp tblogistics.TaobaologisticsordersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
