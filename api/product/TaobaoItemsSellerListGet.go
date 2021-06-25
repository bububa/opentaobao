package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
批量获取商品详细信息 
taobao.items.seller.list.get

批量获取商品详细信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
func TaobaoItemsSellerListGet(clt *core.SDKClient, req *product.TaobaoItemsSellerListGetRequest, session string) (*product.TaobaoItemsSellerListGetResponse, error) {
    var resp product.TaobaoItemsSellerListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
