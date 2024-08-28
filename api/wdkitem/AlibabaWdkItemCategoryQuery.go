package wdkitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemCategoryQuery 类目查询接口
// alibaba.wdk.item.category.query
//
// 类目查询接口
func AlibabaWdkItemCategoryQuery(ctx context.Context, clt *core.SDKClient, req *wdkitem.AlibabaWdkItemCategoryQueryAPIRequest, resp *wdkitem.AlibabaWdkItemCategoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
