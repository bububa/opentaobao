package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
删除SKU 
taobao.item.sku.delete

删除一个sku的数据<br/>需要删除的sku通过属性properties进行匹配查找
*/
func TaobaoItemSkuDelete(clt *core.SDKClient, req *product.TaobaoItemSkuDeleteRequest, session string) (*product.TaobaoItemSkuDeleteResponse, error) {
    var resp product.TaobaoItemSkuDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
