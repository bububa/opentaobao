package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkStockRealQuery 仓内实时库存查询
// alibaba.wdk.stock.real.query
//
// 查询仓内实时库存信息
func AlibabaWdkStockRealQuery(clt *core.SDKClient, req *wdk.AlibabaWdkStockRealQueryAPIRequest, resp *wdk.AlibabaWdkStockRealQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
