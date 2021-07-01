package icbushowcase

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbushowcase"
)

/* 
批量删除橱窗商品 
alibaba.scbp.showcase.deleteproduct

批量删除橱窗商品
*/
func AlibabaScbpShowcaseDeleteproduct(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseDeleteproductAPIRequest, session string) (*icbushowcase.AlibabaScbpShowcaseDeleteproductAPIResponse, error) {
    var resp icbushowcase.AlibabaScbpShowcaseDeleteproductAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
