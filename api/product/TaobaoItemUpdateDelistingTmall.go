package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
taobao.item.update.delisting.tmall 
taobao.item.update.delisting.tmall

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
*/
func TaobaoItemUpdateDelistingTmall(clt *core.SDKClient, req *product.TaobaoItemUpdateDelistingTmallRequest, session string) (*product.TaobaoItemUpdateDelistingTmallResponse, error) {
    var resp product.TaobaoItemUpdateDelistingTmallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
