package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsSellerOrdersGet 商家配送核销订单列表查询
// alibaba.ascp.logistics.seller.orders.get
//
// 商家配送核销订单列表查询
func AlibabaAscpLogisticsSellerOrdersGet(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsSellerOrdersGetAPIRequest, session string) (*tblogistics.AlibabaAscpLogisticsSellerOrdersGetAPIResponse, error) {
	var resp tblogistics.AlibabaAscpLogisticsSellerOrdersGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
