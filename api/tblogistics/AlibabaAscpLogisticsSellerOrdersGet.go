package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticssellerordersget 商家配送核销订单列表查询
// alibaba.ascp.logistics.seller.orders.get
//
// 商家配送核销订单列表查询
func Alibabaascplogisticssellerordersget(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticssellerordersgetAPIRequest, session string) (*tblogistics.AlibabaascplogisticssellerordersgetAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticssellerordersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
