package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCombineskuUpdate 组合商品更新接口
// alibaba.wdk.sku.combinesku.update
//
// 组合商品修改接口
func AlibabaWdkSkuCombineskuUpdate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuCombineskuUpdateAPIRequest, resp *wdk.AlibabaWdkSkuCombineskuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
