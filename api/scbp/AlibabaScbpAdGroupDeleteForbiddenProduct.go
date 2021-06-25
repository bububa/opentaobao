package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
删除屏蔽品 
alibaba.scbp.ad.group.delete.forbidden.product

删除屏蔽品
*/
func AlibabaScbpAdGroupDeleteForbiddenProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupDeleteForbiddenProductRequest, session string) (*scbp.AlibabaScbpAdGroupDeleteForbiddenProductResponse, error) {
    var resp scbp.AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
