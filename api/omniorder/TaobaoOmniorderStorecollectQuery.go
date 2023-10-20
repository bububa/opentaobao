package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstorecollectquery 全渠道门店自提根据核销码查询订单
// taobao.omniorder.storecollect.query
//
// 全渠道门店自提根据核销码查询订单
func Taobaoomniorderstorecollectquery(clt *core.SDKClient, req *omniorder.TaobaoomniorderstorecollectqueryAPIRequest, session string) (*omniorder.TaobaoomniorderstorecollectqueryAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstorecollectqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
