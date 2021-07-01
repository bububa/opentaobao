package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
重命名关键词分组 
alibaba.scbp.tag.rename

重命名关键词分组
*/
func AlibabaScbpTagRename(clt *core.SDKClient, req *scbp.AlibabaScbpTagRenameAPIRequest, session string) (*scbp.AlibabaScbpTagRenameAPIResponse, error) {
    var resp scbp.AlibabaScbpTagRenameAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
