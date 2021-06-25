package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
编辑商品发布页 
taobao.banamadpc.item.edit.render

巴拿马供应商通过此接口获取编辑商品发布页
*/
func TaobaoBanamadpcItemEditRender(clt *core.SDKClient, req *product.TaobaoBanamadpcItemEditRenderRequest, session string) (*product.TaobaoBanamadpcItemEditRenderResponse, error) {
    var resp product.TaobaoBanamadpcItemEditRenderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
