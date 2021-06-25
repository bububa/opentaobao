package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取SKU 
taobao.item.sku.get

获取sku_id所对应的sku数据 
sku_id对应的sku要属于传入的nick对应的卖家
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
func TaobaoItemSkuGet(clt *core.SDKClient, req *product.TaobaoItemSkuGetRequest, session string) (*product.TaobaoItemSkuGetResponse, error) {
    var resp product.TaobaoItemSkuGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
