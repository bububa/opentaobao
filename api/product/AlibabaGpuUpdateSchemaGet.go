package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取产品编辑schema规则的接口 
alibaba.gpu.update.schema.get

获取产品编辑schema规则的接口
*/
func AlibabaGpuUpdateSchemaGet(clt *core.SDKClient, req *product.AlibabaGpuUpdateSchemaGetRequest, session string) (*product.AlibabaGpuUpdateSchemaGetAPIResponse, error) {
    var resp product.AlibabaGpuUpdateSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
