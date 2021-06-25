package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取天猫商品尺码表模板列表 
tmall.item.sizemapping.templates.list

获取所有尺码表模板列表。
*/
func TmallItemSizemappingTemplatesList(clt *core.SDKClient, req *product.TmallItemSizemappingTemplatesListRequest, session string) (*product.TmallItemSizemappingTemplatesListResponse, error) {
    var resp product.TmallItemSizemappingTemplatesListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
