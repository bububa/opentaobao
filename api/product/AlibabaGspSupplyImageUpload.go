package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
gsp图片上传 
alibaba.gsp.supply.image.upload

上传图片至目标海外平台的素材空间
*/
func AlibabaGspSupplyImageUpload(clt *core.SDKClient, req *product.AlibabaGspSupplyImageUploadRequest, session string) (*product.AlibabaGspSupplyImageUploadAPIResponse, error) {
    var resp product.AlibabaGspSupplyImageUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
