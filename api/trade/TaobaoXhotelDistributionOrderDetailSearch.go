package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaoxhoteldistributionorderdetailsearch 分销渠道订单详情查询
// taobao.xhotel.distribution.order.detail.search
//
// 该接口用于提供酒店分销渠道的订单详情查询
func Taobaoxhoteldistributionorderdetailsearch(clt *core.SDKClient, req *trade.TaobaoxhoteldistributionorderdetailsearchAPIRequest, session string) (*trade.TaobaoxhoteldistributionorderdetailsearchAPIResponse, error) {
	var resp trade.TaobaoxhoteldistributionorderdetailsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
