package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
按类目查询spu接口 
alibaba.gpu.schema.catsearch

按类目查询spu的schema接口
*/
func AlibabaGpuSchemaCatsearch(clt *core.SDKClient, req *product.AlibabaGpuSchemaCatsearchAPIRequest, session string) (*product.AlibabaGpuSchemaCatsearchAPIResponse, error) {
    var resp product.AlibabaGpuSchemaCatsearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
