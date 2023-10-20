package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsSkuCreate 商品同步
// taobao.wlb.wms.sku.create
//
// 商品同步
func TaobaoWlbWmsSkuCreate(clt *core.SDKClient, req *wms.TaobaoWlbWmsSkuCreateAPIRequest, resp *wms.TaobaoWlbWmsSkuCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
