package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
外贸直通车批量删除关键词 
alibaba.scbp.ad.keyword.batchdelete

外贸直通车批量删除关键词
*/
func AlibabaScbpAdKeywordBatchdelete(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordBatchdeleteRequest, session string) (*scbp.AlibabaScbpAdKeywordBatchdeleteAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordBatchdeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
