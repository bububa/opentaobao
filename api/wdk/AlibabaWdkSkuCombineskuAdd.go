package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCombineskuAdd 组合商品新增接口
// alibaba.wdk.sku.combinesku.add
//
// 组合商品新增接口
func AlibabaWdkSkuCombineskuAdd(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuCombineskuAddAPIRequest, resp *wdk.AlibabaWdkSkuCombineskuAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
