package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsSkuUpdate 商品信息的更新
// taobao.wlb.wms.sku.update
//
// 商品信息的更新
func TaobaoWlbWmsSkuUpdate(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsSkuUpdateAPIRequest, resp *wms.TaobaoWlbWmsSkuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
