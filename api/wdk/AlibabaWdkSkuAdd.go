package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuAdd 新增商品
// alibaba.wdk.sku.add
//
// 创建RT门店商品或DC商品
func AlibabaWdkSkuAdd(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuAddAPIRequest, resp *wdk.AlibabaWdkSkuAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
