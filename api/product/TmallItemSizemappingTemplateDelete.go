package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
删除天猫商品尺码表模板 
tmall.item.sizemapping.template.delete

删除天猫商品尺码表模板
*/
func TmallItemSizemappingTemplateDelete(clt *core.SDKClient, req *product.TmallItemSizemappingTemplateDeleteRequest, session string) (*product.TmallItemSizemappingTemplateDeleteResponse, error) {
    var resp product.TmallItemSizemappingTemplateDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
