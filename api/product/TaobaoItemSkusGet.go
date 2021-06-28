package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
根据商品ID列表获取SKU信息 
taobao.item.skus.get

* 获取多个商品下的所有sku
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
func TaobaoItemSkusGet(clt *core.SDKClient, req *product.TaobaoItemSkusGetRequest, session string) (*product.TaobaoItemSkusGetAPIResponse, error) {
    var resp product.TaobaoItemSkusGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
