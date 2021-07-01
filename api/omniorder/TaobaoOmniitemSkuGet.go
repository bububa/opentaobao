package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoOmniitemSkuGet
获取全渠道门店商品sku
taobao.omniitem.sku.get

通过skuId或者skuOutId查询全渠道门店商品sku信息 */
func TaobaoOmniitemSkuGet(clt *core.SDKClient, req *omniorder.TaobaoOmniitemSkuGetAPIRequest, session string) (*omniorder.TaobaoOmniitemSkuGetAPIResponse, error) {
	var resp omniorder.TaobaoOmniitemSkuGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
