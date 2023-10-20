package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Alibabaretailcommissionorderquery 分销订单查询
// alibaba.retail.commission.order.query
//
// 查询商家的分销订单
func Alibabaretailcommissionorderquery(clt *core.SDKClient, req *omniorder.AlibabaretailcommissionorderqueryAPIRequest, session string) (*omniorder.AlibabaretailcommissionorderqueryAPIResponse, error) {
	var resp omniorder.AlibabaretailcommissionorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
