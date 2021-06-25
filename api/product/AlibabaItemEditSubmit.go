package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品编辑提交schema信息 
alibaba.item.edit.submit

商品编辑提交schema信息
*/
func AlibabaItemEditSubmit(clt *core.SDKClient, req *product.AlibabaItemEditSubmitRequest, session string) (*product.AlibabaItemEditSubmitResponse, error) {
    var resp product.AlibabaItemEditSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
