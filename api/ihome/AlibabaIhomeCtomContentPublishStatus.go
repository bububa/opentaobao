package ihome

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ihome"
)

/* 
实拍图发布审核状态查询API 
alibaba.ihome.ctom.content.publish.status

实拍图发布审核状态查询API
*/
func AlibabaIhomeCtomContentPublishStatus(clt *core.SDKClient, req *ihome.AlibabaIhomeCtomContentPublishStatusAPIRequest, session string) (*ihome.AlibabaIhomeCtomContentPublishStatusAPIResponse, error) {
    var resp ihome.AlibabaIhomeCtomContentPublishStatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
