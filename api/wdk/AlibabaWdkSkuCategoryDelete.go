package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCategoryDelete 商家类目删除接口
// alibaba.wdk.sku.category.delete
//
// 商家类目删除接口
func AlibabaWdkSkuCategoryDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryDeleteAPIRequest, resp *wdk.AlibabaWdkSkuCategoryDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
