package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
创建初始模板 
alibaba.wholesale.shippingline.template.init

创建默认的几种运费模板
*/
func AlibabaWholesaleShippinglineTemplateInit(clt *core.SDKClient, req *product.AlibabaWholesaleShippinglineTemplateInitRequest, session string) (*product.AlibabaWholesaleShippinglineTemplateInitResponse, error) {
    var resp product.AlibabaWholesaleShippinglineTemplateInitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
