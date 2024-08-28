package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBmPaiyangStockQuery 派样商品门店库存查询接口
// alibaba.wdk.bm.paiyang.stock.query
//
// 淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。
func AlibabaWdkBmPaiyangStockQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkBmPaiyangStockQueryAPIRequest, resp *wdk.AlibabaWdkBmPaiyangStockQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
