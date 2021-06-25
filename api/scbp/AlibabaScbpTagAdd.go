package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
创建关键词分组 
alibaba.scbp.tag.add

创建关键词分组
*/
func AlibabaScbpTagAdd(clt *core.SDKClient, req *scbp.AlibabaScbpTagAddRequest, session string) (*scbp.AlibabaScbpTagAddResponse, error) {
    var resp scbp.AlibabaScbpTagAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
