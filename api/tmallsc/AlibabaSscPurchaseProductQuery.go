package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaSscPurchaseProductQuery 查询已采购的服务产品
// alibaba.ssc.purchase.product.query
//
// 查询已采购的服务产品
func AlibabaSscPurchaseProductQuery(clt *core.SDKClient, req *tmallsc.AlibabaSscPurchaseProductQueryAPIRequest, resp *tmallsc.AlibabaSscPurchaseProductQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
