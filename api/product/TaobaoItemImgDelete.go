package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
删除商品图片 
taobao.item.img.delete

删除商品图片
*/
func TaobaoItemImgDelete(clt *core.SDKClient, req *product.TaobaoItemImgDeleteRequest, session string) (*product.TaobaoItemImgDeleteResponse, error) {
    var resp product.TaobaoItemImgDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
