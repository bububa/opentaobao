package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
修改关键词所属分组 
alibaba.scbp.ad.keyword.tag.update

修改关键词所属分组
*/
func AlibabaScbpAdKeywordTagUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordTagUpdateAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordTagUpdateAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordTagUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
