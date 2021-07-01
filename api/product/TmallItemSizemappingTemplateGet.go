package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取天猫商品尺码表模板 
tmall.item.sizemapping.template.get

获取天猫商品尺码表模板
*/
func TmallItemSizemappingTemplateGet(clt *core.SDKClient, req *product.TmallItemSizemappingTemplateGetAPIRequest, session string) (*product.TmallItemSizemappingTemplateGetAPIResponse, error) {
    var resp product.TmallItemSizemappingTemplateGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
