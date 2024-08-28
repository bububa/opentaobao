package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsSkuGet 商品信息查询
// taobao.wlb.wms.sku.get
//
// 商品信息查询
func TaobaoWlbWmsSkuGet(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsSkuGetAPIRequest, resp *wms.TaobaoWlbWmsSkuGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
