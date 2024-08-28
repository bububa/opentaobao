package wdkitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemStoreskuQuery 门店商品信息查询
// alibaba.wdk.item.storesku.query
//
// 门店商品信息查询
func AlibabaWdkItemStoreskuQuery(ctx context.Context, clt *core.SDKClient, req *wdkitem.AlibabaWdkItemStoreskuQueryAPIRequest, resp *wdkitem.AlibabaWdkItemStoreskuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
