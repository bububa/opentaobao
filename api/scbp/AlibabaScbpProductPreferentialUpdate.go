package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
设置P4P产品优先推广状态 
alibaba.scbp.product.preferential.update

设置P4P产品优先推广状态
*/
func AlibabaScbpProductPreferentialUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpProductPreferentialUpdateAPIRequest, session string) (*scbp.AlibabaScbpProductPreferentialUpdateAPIResponse, error) {
    var resp scbp.AlibabaScbpProductPreferentialUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
