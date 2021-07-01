package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取一个产品的信息 
taobao.product.get

天猫商家发布商品时，查询关联产品信息时使用，非商品查询接口。商品查询接口：taobao.item.seller.get<br>
两种方式查看一个产品详细信息: 
传入product_id来查询；传入cid和props来查询
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
func TaobaoProductGet(clt *core.SDKClient, req *product.TaobaoProductGetAPIRequest, session string) (*product.TaobaoProductGetAPIResponse, error) {
    var resp product.TaobaoProductGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
