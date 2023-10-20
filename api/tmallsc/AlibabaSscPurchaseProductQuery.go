package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaSscPurchaseProductQuery 查询已采购的服务产品
// alibaba.ssc.purchase.product.query
//
// 查询已采购的服务产品
func AlibabaSscPurchaseProductQuery(clt *core.SDKClient, req *tmallsc.AlibabaSscPurchaseProductQueryAPIRequest, session string) (*tmallsc.AlibabaSscPurchaseProductQueryAPIResponse, error) {
	var resp tmallsc.AlibabaSscPurchaseProductQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
