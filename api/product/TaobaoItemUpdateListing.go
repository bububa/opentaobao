package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
一口价商品上架 
taobao.item.update.listing

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户
*/
func TaobaoItemUpdateListing(clt *core.SDKClient, req *product.TaobaoItemUpdateListingRequest, session string) (*product.TaobaoItemUpdateListingResponse, error) {
    var resp product.TaobaoItemUpdateListingAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
