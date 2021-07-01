package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询所有分组 
alibaba.scbp.tag.list

查询所有分组
*/
func AlibabaScbpTagList(clt *core.SDKClient, req *scbp.AlibabaScbpTagListAPIRequest, session string) (*scbp.AlibabaScbpTagListAPIResponse, error) {
    var resp scbp.AlibabaScbpTagListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
