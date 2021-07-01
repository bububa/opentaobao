package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取产品发布规则接口 
alibaba.gpu.add.schema.get

获取产品发布规则接口
*/
func AlibabaGpuAddSchemaGet(clt *core.SDKClient, req *product.AlibabaGpuAddSchemaGetAPIRequest, session string) (*product.AlibabaGpuAddSchemaGetAPIResponse, error) {
    var resp product.AlibabaGpuAddSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
