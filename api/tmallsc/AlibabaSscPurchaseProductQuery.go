package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabasscpurchaseproductquery 查询已采购的服务产品
// alibaba.ssc.purchase.product.query
//
// 查询已采购的服务产品
func Alibabasscpurchaseproductquery(clt *core.SDKClient, req *tmallsc.AlibabasscpurchaseproductqueryAPIRequest, session string) (*tmallsc.AlibabasscpurchaseproductqueryAPIResponse, error) {
	var resp tmallsc.AlibabasscpurchaseproductqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
