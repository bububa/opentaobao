package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
修改P4P产品推广状态 
alibaba.scbp.product.status.update

修改P4P产品推广状态
*/
func AlibabaScbpProductStatusUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpProductStatusUpdateRequest, session string) (*scbp.AlibabaScbpProductStatusUpdateResponse, error) {
    var resp scbp.AlibabaScbpProductStatusUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
