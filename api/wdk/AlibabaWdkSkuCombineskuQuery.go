package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCombineskuQuery 组合商品查询接口
// alibaba.wdk.sku.combinesku.query
//
// 查询组合商品接口
func AlibabaWdkSkuCombineskuQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuCombineskuQueryAPIRequest, resp *wdk.AlibabaWdkSkuCombineskuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
