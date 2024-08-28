package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuWarehouseskuQuery 仓商品查询接口(指定商品编码)
// alibaba.wdk.sku.warehousesku.query
//
// 提供指定仓商品编码查询
func AlibabaWdkSkuWarehouseskuQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuWarehouseskuQueryAPIRequest, resp *wdk.AlibabaWdkSkuWarehouseskuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
