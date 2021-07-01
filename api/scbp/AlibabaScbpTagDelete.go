package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
删除关键词分组 
alibaba.scbp.tag.delete

删除关键词分组
*/
func AlibabaScbpTagDelete(clt *core.SDKClient, req *scbp.AlibabaScbpTagDeleteAPIRequest, session string) (*scbp.AlibabaScbpTagDeleteAPIResponse, error) {
    var resp scbp.AlibabaScbpTagDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
