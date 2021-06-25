package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取单个商品详细信息 
taobao.item.seller.get

获取单个商品的全部信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
func TaobaoItemSellerGet(clt *core.SDKClient, req *product.TaobaoItemSellerGetRequest, session string) (*product.TaobaoItemSellerGetResponse, error) {
    var resp product.TaobaoItemSellerGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
