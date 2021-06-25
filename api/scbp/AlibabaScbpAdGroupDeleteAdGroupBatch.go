package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
删除推广单元 
alibaba.scbp.ad.group.delete.ad.group.batch

删除推广单元
*/
func AlibabaScbpAdGroupDeleteAdGroupBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupDeleteAdGroupBatchRequest, session string) (*scbp.AlibabaScbpAdGroupDeleteAdGroupBatchResponse, error) {
    var resp scbp.AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
