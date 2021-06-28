package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品删除 
alibaba.item.operate.delete

商品删除
*/
func AlibabaItemOperateDelete(clt *core.SDKClient, req *product.AlibabaItemOperateDeleteRequest, session string) (*product.AlibabaItemOperateDeleteAPIResponse, error) {
    var resp product.AlibabaItemOperateDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
