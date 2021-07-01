package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品下架 
taobao.item.update.delisting

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
*/
func TaobaoItemUpdateDelisting(clt *core.SDKClient, req *product.TaobaoItemUpdateDelistingAPIRequest, session string) (*product.TaobaoItemUpdateDelistingAPIResponse, error) {
    var resp product.TaobaoItemUpdateDelistingAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
