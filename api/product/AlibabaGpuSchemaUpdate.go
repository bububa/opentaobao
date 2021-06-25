package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
产品更新接口 
alibaba.gpu.schema.update

产品更新接口
*/
func AlibabaGpuSchemaUpdate(clt *core.SDKClient, req *product.AlibabaGpuSchemaUpdateRequest, session string) (*product.AlibabaGpuSchemaUpdateResponse, error) {
    var resp product.AlibabaGpuSchemaUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
