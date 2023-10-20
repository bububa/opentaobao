package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniitemSkuGet 获取全渠道门店商品sku
// taobao.omniitem.sku.get
//
// 通过skuId或者skuOutId查询全渠道门店商品sku信息
func TaobaoOmniitemSkuGet(clt *core.SDKClient, req *omniorder.TaobaoOmniitemSkuGetAPIRequest, resp *omniorder.TaobaoOmniitemSkuGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
