package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCategoryQuery 商家类目获取接口
// alibaba.wdk.sku.category.query
//
// 商家类目获取接口
func AlibabaWdkSkuCategoryQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryQueryAPIRequest, resp *wdk.AlibabaWdkSkuCategoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
