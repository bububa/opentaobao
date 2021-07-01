package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
关键词启动暂停推广 
alibaba.scbp.ad.keyword.status.update

关键词启动暂停推广
*/
func AlibabaScbpAdKeywordStatusUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordStatusUpdateAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordStatusUpdateAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordStatusUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
